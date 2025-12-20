// internal/notification-service/adapter/inbound/kafka/signup_consumer.go
package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/notification-service/application/port/inbound"
	"github.com/anfastk/MERGESPACE/shared/avro"
	"github.com/riferrei/srclient"
	"github.com/twmb/franz-go/pkg/kgo"
)

type SignupConsumer struct {
	client  *kgo.Client
	handler inbound.OTPEventHandler
	sr      *srclient.SchemaRegistryClient
}

func NewSignupConsumer(
	brokers []string,
	group string,
	topic string,
	srURL string,
	handler inbound.OTPEventHandler,
) (*SignupConsumer, error) {

	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(group),
		kgo.ConsumeTopics(topic),
	)
	if err != nil {
		return nil, err
	}

	return &SignupConsumer{
		client:  client,
		handler: handler,
		sr:      srclient.CreateSchemaRegistryClient(srURL),
	}, nil
}

func (c *SignupConsumer) Start(ctx context.Context) {
	for {
		fetches := c.client.PollFetches(ctx)
		if fetches.IsClientClosed() {
			return
		}

		fetches.EachError(func(topic string, partition int32, err error) {
			log.Printf("Kafka error %s[%d]: %v", topic, partition, err)
		})

		fetches.EachRecord(func(rec *kgo.Record) {
			eventMap, err := avro.Decode(c.sr, rec.Value)
			if err != nil {
				log.Printf("Avro decode failed: %v", err)
				return
			}

			fmt.Println("eventMap",eventMap)

			event := dto.SignupOTPEvent{
				Email:     eventMap["email"].(string),
				OTP:       eventMap["otp"].(string),
				CreatedAt: int64(eventMap["created_at"].(int64)),
			}
			fmt.Println("EVENT :",event)

			if err := c.handler.HandleSignupOTP(ctx, event); err != nil {
				log.Printf("Handler failed: %v", err)
				return
			}
		})
	}
}

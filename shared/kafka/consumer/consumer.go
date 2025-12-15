package consumer

import (
	"context"
	"log"

	"github.com/anfastk/MERGESPACE/shared/avro"
	"github.com/riferrei/srclient"
	"github.com/twmb/franz-go/pkg/kgo"
)

type HandlerFunc func(context.Context, map[string]interface{}) error

type Consumer struct {
	client *kgo.Client
	sr     *srclient.SchemaRegistryClient
	handle HandlerFunc
}

func New(brokers []string, srURL, group string, topics []string, handler HandlerFunc) (*Consumer, error) {

	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(group),
		kgo.ConsumeTopics(topics...),
	)
	if err != nil {
		return nil, err
	}

	sr := srclient.CreateSchemaRegistryClient(srURL)

	return &Consumer{client, sr, handler}, nil
}

func (c *Consumer) Run(ctx context.Context) {
	for {
		fetches := c.client.PollFetches(ctx)
		if fetches.IsClientClosed() {
			return
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()

			ev, err := avro.Decode(c.sr, record.Value)
			if err != nil {
				log.Println("decode error:", err)
				continue
			}

			if err := c.handle(ctx, ev); err != nil {
				log.Println("handler error:", err)
				continue
			}

			c.client.CommitRecords(ctx, record)
		}
	}
}

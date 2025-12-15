package kafka

import (
	"context"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/outbound"
	"github.com/anfastk/MERGESPACE/shared/kafka/producer"
)

type SignupEventProducer struct {
	prod   *producer.Producer
	schema string
}

func NewSignupEventProducer(prod *producer.Producer) outbound.OTPEventPublisher {
	return &SignupEventProducer{
		prod: prod,
	}
}

func (p *SignupEventProducer) PublishOTPEvent(ctx context.Context, event dto.SignupOTPEvent) error {
	avroEvt := toAvroEvent(event)
	return p.prod.Publish(ctx, []byte(avroEvt.Email), avroEvt)

}

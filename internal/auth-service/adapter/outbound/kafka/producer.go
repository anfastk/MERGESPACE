package kafka

import (
	"github.com/anfastk/MERGESPACE/internal/auth-service/adapter/outbound/kafka/avro"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
)

func toAvroEvent(e dto.SignupOTPEvent) avro.SignupOTPEventAvro {
	return avro.SignupOTPEventAvro{
		SignupSessionID: e.SignupSessionID,
		Email:           e.Email,
		OTP:             e.OTP,
		CreatedAt:       e.CreatedAt,
	}
}

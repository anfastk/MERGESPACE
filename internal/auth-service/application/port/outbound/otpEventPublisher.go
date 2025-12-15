package outbound

import (
	"context"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
)

type OTPEventPublisher interface {
	PublishOTPEvent(ctx context.Context,  event dto.SignupOTPEvent) error
}

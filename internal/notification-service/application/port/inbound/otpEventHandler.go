package inbound

import (
	"context"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
)

type OTPEventHandler interface {
	HandleSignupOTP(ctx context.Context, event dto.SignupOTPEvent) error
}

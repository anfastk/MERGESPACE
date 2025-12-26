// internal/notification-service/application/service/otp_handler.go
package service

import (
	"context"
	"log"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
)

type OTPHandler struct {
}

func NewOTPHandler() *OTPHandler {
	return &OTPHandler{}
}

func (h *OTPHandler) HandleSignupOTP(ctx context.Context, event dto.SignupOTPEvent) error {
	log.Printf("Sending OTP %s to %s", event.OTP, event.Email)

	// TODO:
	// - call email provider
	// - retry logic
	// - idempotency

	return nil
}

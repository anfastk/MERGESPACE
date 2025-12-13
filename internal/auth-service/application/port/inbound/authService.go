package inbound

import (
	"context"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
)

type AuthUseCase interface {
	InitiateSignup(ctx context.Context, req dto.InitiateSignUpRequest) (*dto.InitiateSignUpResponse, error)
}

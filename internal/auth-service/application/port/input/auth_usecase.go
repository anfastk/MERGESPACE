package input

import (
	"context"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
)

type AuthUseCase interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error)
}

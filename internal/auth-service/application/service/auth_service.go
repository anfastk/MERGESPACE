package service

import (
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/output"
)

type AuthService struct {
	userRepo    output.UserRepository
	sessionRepo output.SessionRepository
}

func NewAuthService(uRepo output.UserRepository, sRepo output.SessionRepository) *AuthService {
	return &AuthService{
		userRepo:    uRepo,
		sessionRepo: sRepo,
	}
}

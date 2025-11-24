package service

import (
	"context"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/apperrors"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, apperrors.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, apperrors.ErrInvalidCredentials
	}

	return s.generateSession(ctx, user, req.IPAddress, req.UserAgent)
}

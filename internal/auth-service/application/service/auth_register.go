package service

import (
	"context"
	"time"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/apperrors"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error) {
	existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, apperrors.ErrUserAlreadyExists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}

	newUser := &entity.User{
		Email:         req.Email,
		Username:      req.Username,
		PasswordHash:  string(hashed),
		AccountStatus: "active",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.userRepo.Save(ctx, newUser); err != nil {
		return nil, apperrors.ErrInternalServer
	}
	return s.generateSession(ctx, newUser, "registration_ip", "registration_ua")
}

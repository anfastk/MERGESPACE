package service

import (
	"context"
	"time"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/apperrors"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
	"github.com/google/uuid"
)

func (s *AuthService) generateSession(ctx context.Context, user *entity.User, ip, ua string) (*dto.AuthResponse, error) {
	accessToken := uuid.New().String()
	refreshToken := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)

	session := &entity.Session{
		SessionID:    uuid.New().String(),
		UserID:       uint64(user.ID),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now(),
		IPAddress:    ip,
		UserAgent:    ua,
	}
	if err := s.sessionRepo.CreateSession(ctx, session); err != nil {
		return nil, apperrors.ErrInternalServer
	}
	return &dto.AuthResponse{
		UserID:       uint64(user.ID),
		Username:     user.Username,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}

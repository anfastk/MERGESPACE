package dto

import "github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"

type OAuthLoginRequest struct {
	Provider    entity.AuthProvider
	AccessToken string
}

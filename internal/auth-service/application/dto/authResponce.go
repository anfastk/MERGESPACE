package dto

import "github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"

type AuthResponse struct {
	User   *entity.User
	Tokens *AuthTokens
}

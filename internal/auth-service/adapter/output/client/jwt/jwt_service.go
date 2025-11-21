package jwt

import (
	"time"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/output"
	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenService struct {
	secret string
}

func NewJwtTokenService(secret string) output.TokenService {
	return &JwtTokenService{secret: secret}
}

func (j *JwtTokenService) GenerateTokens(userID int64, email string) (string, string, error) {

	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := access.SignedString([]byte(j.secret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refresh.SignedString([]byte(j.secret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

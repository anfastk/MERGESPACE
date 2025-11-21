package di

import (
	authInput "github.com/anfastk/MERGESPACE/internal/auth-service/adapter/input/http"
	jwtClient "github.com/anfastk/MERGESPACE/internal/auth-service/adapter/output/client/jwt"
	authRepo "github.com/anfastk/MERGESPACE/internal/auth-service/adapter/output/persistence/postgres"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/service"
	"gorm.io/gorm"
)

func InitializeAuth(db *gorm.DB) *authInput.AuthHandler {

	userRepo := authRepo.NewUserRepository(db)
	tokenService := jwtClient.NewJwtTokenService("SECRET_KEY")

	authService := service.NewAuthService(userRepo, tokenService)

	return authInput.NewAuthHandler(authService)
}

package di

/* import (
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/service"
	"github.com/anfastk/MERGESPACE/internal/auth-service/infrastructure/config"
	"github.com/anfastk/MERGESPACE/internal/auth-service/infrastructure/server"
	"gorm.io/gorm"
)

func InitializeApp(db *gorm.DB, cfg *config.Config) (*server.Server, error) {
	authRepo := postgres.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo, authRepo)
	authHandler := http.NewAuthHandler(authService)

	return server.NewServer(authHandler, cfg.ServerPort), nil
}
 */
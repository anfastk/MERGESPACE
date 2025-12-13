package outbound

import (
	"context"

	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
}

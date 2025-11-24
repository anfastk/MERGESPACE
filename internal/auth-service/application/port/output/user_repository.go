package output

import (
	"context"

	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	IDExists(ctx context.Context, id uint64) (bool, error)
}

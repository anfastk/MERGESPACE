package outbound

import (
	"context"

	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
)

type PendingSignupRepository interface {
	Save(ctx context.Context, signup *entity.PendingSignup) error
}

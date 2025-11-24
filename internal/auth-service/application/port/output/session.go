package output

import (
	"context"

	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session *entity.Session) error
}

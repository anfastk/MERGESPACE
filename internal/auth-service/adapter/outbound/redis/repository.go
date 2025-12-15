package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/outbound"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
	"github.com/redis/go-redis/v9"
)

const pendingSignupPrefix = "signup:pending:"

type PendingSignupRedisRepository struct {
	client *redis.Client
}

var _ outbound.PendingSignupRepository = (*PendingSignupRedisRepository)(nil)

func NewPendingSignupRepository(client *redis.Client) outbound.PendingSignupRepository {
	return &PendingSignupRedisRepository{client: client}
}

func (r *PendingSignupRedisRepository) key(id entity.PendingSignupID) string {
	return fmt.Sprintf("%s%s", pendingSignupPrefix, id)
}

func (r *PendingSignupRedisRepository) Save(ctx context.Context, signup *entity.PendingSignup) error {
	model := pendingSignupModel{
		ID:           string(signup.ID),
		FirstName:    signup.FirstName,
		LastName:     signup.LastName,
		Email:        signup.Email,
		Username:     signup.Username,
		PasswordHash: signup.PasswordHash,
		OTP:          signup.OTP,
		Attempts:     signup.Attempts,
		ResendCount:  signup.ResendCount,
		CreatedAt:    signup.CreatedAt,
		ExpiresAt:    signup.ExpiresAt,
	}

	data, err := json.Marshal(model)
	if err != nil {
		return err
	}

	ttl := time.Until(signup.ExpiresAt)
	if ttl <= 0 {
		ttl = time.Minute
	}

	return r.client.Set(ctx, r.key(signup.ID), data, ttl).Err()
}

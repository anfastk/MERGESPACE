package algorithm

import (
	"context"

	"github.com/anfastk/MERGESPACE/shared/ratelimiter/limiter/backend"
)

type TokenBucket struct {
	store backend.Store
}

func NewTokenBucket(store backend.Store) *TokenBucket {
	return &TokenBucket{store: store}
}

func (t *TokenBucket) Name() string {
	return "token_bucket"
}

func (t *TokenBucket) Allow(ctx context.Context, key string, limit int, windowSec int) (bool, int, error) {
	return t.store.Eval(
		ctx,
		"token_bucket",
		[]string{key},
		limit,
		windowSec,
	)
}

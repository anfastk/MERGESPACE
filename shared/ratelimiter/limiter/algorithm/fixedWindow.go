package algorithm

import (
	"context"

	"github.com/anfastk/MERGESPACE/shared/ratelimiter/limiter/backend"
)

type FixedWindow struct {
	store backend.Store
}

func NewFixedWindow(store backend.Store) *FixedWindow {
	return &FixedWindow{store: store}
}

func (f *FixedWindow) Name() string {
	return "fixed_window"
}

func (f *FixedWindow) Allow(ctx context.Context, key string, limit int, windowSec int) (bool, int, error) {
	return f.store.Eval(
		ctx,
		"fixed_window",
		[]string{key},
		limit,
		windowSec,
	)
}

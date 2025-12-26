package algorithm

import "context"

type Algorithm interface {
	Name() string
	Allow(ctx context.Context, key string, limit int, windowSec int) (bool, int, error)
}

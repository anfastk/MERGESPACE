package backend

import "context"

type Store interface {
	Eval(ctx context.Context, script string, keys []string, args ...any) (allowed bool, retryAfter int, err error)
}

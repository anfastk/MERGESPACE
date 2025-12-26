package limiter

import (
	"context"
	"errors"
)

type Algorithm interface {
	Allow(ctx context.Context, key string, limit int, window int) (bool, int, error)
}

type Limiter struct {
	algorithms map[string]Algorithm
}

func NewLimiter(algos []Algorithm) *Limiter {
	m := make(map[string]Algorithm)
	for _, a := range algos {
		name := a.(interface{ Name() string }).Name()
		m[name] = a
	}
	return &Limiter{algorithms: m}
}

func (l *Limiter) Allow(ctx context.Context, rule Rule, identifier string) (bool, int, error) {

	algo, ok := l.algorithms[rule.Algo]
	if !ok {
		return false, 0, errors.New("unknown rate limit algorithm")
	}
	
	key := rule.KeyPrefix + ":" + identifier

	allowed, retry, err := algo.Allow(ctx, key, rule.Limit, rule.WindowSec)

	if err != nil {
		if rule.FailOpen {
			return true, 0, nil
		}
		return false, 0, err
	}

	if !allowed {
		return false, retry, ErrRateLimited
	}

	return true, 0, nil
}

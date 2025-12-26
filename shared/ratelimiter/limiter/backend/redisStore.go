package backend

import (
	"context"
	"errors"
	"strings"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	rdb     *redis.Client
	scripts map[string]string 
}

func NewRedisStore(rdb *redis.Client, scripts map[string]string) *RedisStore {
	return &RedisStore{
		rdb:     rdb,
		scripts: scripts,
	}
}

func (s *RedisStore) Eval(ctx context.Context, script string, keys []string, args ...any) (bool, int, error) {

	sha, ok := s.scripts[script]
	if !ok {
		return false, 0, errors.New("script not registered")
	}

	res, err := s.rdb.EvalSha(ctx, sha, keys, args...).Result()
	if err != nil {
		if strings.Contains(err.Error(), "NOSCRIPT") {
			return false, 0, err
		}
		return false, 0, err
	}

	arr, ok := res.([]interface{})
	if !ok || len(arr) != 2 {
		return false, 0, errors.New("invalid redis response")
	}

	return arr[0].(int64) == 1, int(arr[1].(int64)), nil
}

func LoadScripts(ctx context.Context, rdb *redis.Client) (map[string]string, error) {
	scripts := make(map[string]string)

	tbSHA, err := rdb.ScriptLoad(ctx, tokenBucketScript).Result()
	if err != nil {
		return nil, err
	}
	scripts["token_bucket"] = tbSHA

	fwSHA, err := rdb.ScriptLoad(ctx, fixedWindowScript).Result()
	if err != nil {
		return nil, err
	}
	scripts["fixed_window"] = fwSHA

	return scripts, nil
}

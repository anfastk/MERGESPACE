package limiter

import "errors"

var ErrRateLimited = errors.New("rate limit exceeded")

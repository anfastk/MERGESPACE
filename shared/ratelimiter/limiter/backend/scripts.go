package backend

import _ "embed"

//go:embed scripts/token_bucket.lua
var tokenBucketScript string

//go:embed scripts/fixed_window.lua
var fixedWindowScript string

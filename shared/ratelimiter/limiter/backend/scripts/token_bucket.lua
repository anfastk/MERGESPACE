-- KEYS[1] = key
-- ARGV[1] = capacity
-- ARGV[2] = window_sec

local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local window = tonumber(ARGV[2])

local rate = capacity / window

local data = redis.call("HMGET", key, "tokens", "ts")
local tokens = tonumber(data[1])
local ts = tonumber(data[2])

local now = redis.call("TIME")[1]

if not tokens then
	tokens = capacity
	ts = now
end

local delta = math.max(0, now - ts)
tokens = math.min(capacity, tokens + delta * rate)

if tokens < 1 then
	local retry = math.ceil((1 - tokens) / rate)
	redis.call("HMSET", key, "tokens", tokens, "ts", now)
	redis.call("EXPIRE", key, window)
	return {0, retry}
end

tokens = tokens - 1
redis.call("HMSET", key, "tokens", tokens, "ts", now)
redis.call("EXPIRE", key, window)

return {1, 0}

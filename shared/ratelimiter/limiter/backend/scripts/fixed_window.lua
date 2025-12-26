-- KEYS[1] = key
-- ARGV[1] = limit
-- ARGV[2] = window_sec

local key = KEYS[1]
local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])

local count = redis.call("INCR", key)
if count == 1 then
	redis.call("EXPIRE", key, window)
end

if count > limit then
	local ttl = redis.call("TTL", key)
	return {0, ttl}
end

return {1, 0}

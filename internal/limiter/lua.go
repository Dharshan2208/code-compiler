package limiter

import "github.com/redis/go-redis/v9"

// allows script imeplemt a distributed token bucket rate limiter
// it uses redis hash structure
//   tokens       - available tokens in the bucket.
//   last_refill  - timestamp (in milliseconds) when the bucket was last updated.
//
// ARGV:
//   ARGV[1] - Bucket capacity (maximum number of tokens).
//   ARGV[2] - Refill rate (tokens added per second).
//   ARGV[3] - Current timestamp in milliseconds.
//   ARGV[4] - TTL (expiration time in seconds) for the Redis key.

// Returns:
//   1 -> Request is allowed.
//   0 -> Request is rate limited.

var allowScript = redis.NewScript(`
	local capacity = tonumber(ARGV[1])
	local refillRate = tonumber(ARGV[2])
	local now = tonumber(ARGV[3])
	local ttl = tonumber(ARGV[4])

	-- Fetch the current bucket state from Redis.
	-- "tokens"      = remaining tokens in the bucket.
	-- "last_refill" = last timestamp when refill was calculated.
	local values = redis.call(
		"HMGET",
		KEYS[1],
		"tokens",
		"last_refill"
	)

	local tokens = tonumber(values[1])
	local lastRefill = tonumber(values[2])

	-- Initialize new full bucket if it does exists 
	if not tokens then
		tokens = capacity
		lastRefill = now
	end

	-- Refill tokens
	local elapsed = (now - lastRefill) / 1000.0
	tokens = math.min(
		capacity,
		tokens + (elapsed * refillRate)
	)

	-- Try to consume one token
	local allowed = 0

	if tokens >= 1 then
		tokens = tokens - 1
		allowed = 1
	end

	-- Save updated state atomically done
	redis.call(
		"HMSET",
		KEYS[1],
		"tokens",
		tokens,
		"last_refill",
		now
	)

	-- Automcatically remove inactive buckets after ttl
	redis.call("EXPIRE", KEYS[1], ttl)

	return allowed
`)

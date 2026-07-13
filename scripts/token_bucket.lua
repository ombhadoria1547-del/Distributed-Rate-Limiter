-- Token Bucket Rate Limiter
-- Executes atomically inside Redis

local now = tonumber(ARGV[1])
local capacity = tonumber(ARGV[2])
local refillRate = tonumber(ARGV[3])

local tokens = redis.call("HGET", KEYS[1], "tokens")
local lastRefill = redis.call("HGET", KEYS[1], "last_refill")

if tokens == false then
    tokens = capacity
    lastRefill = now
end

tokens = tonumber(tokens)
lastRefill = tonumber(lastRefill)

local elapsed = now - lastRefill

tokens = tokens + (elapsed * refillRate)
tokens = math.min(tokens, capacity)

lastRefill = now

local allowed = 0

if tokens >= 1 then
    tokens = tokens - 1
    allowed = 1
end

redis.call(
    "HSET",
    KEYS[1],
    "tokens", tokens,
    "last_refill", lastRefill
)

local retryAfter = 0

if allowed == 0 then
    retryAfter = math.ceil((1 - tokens) / refillRate)
end

return {
    allowed,
    math.floor(tokens),
    retryAfter
}
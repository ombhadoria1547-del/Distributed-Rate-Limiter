local windowKey = KEYS[1]

local currentTime = tonumber(ARGV[1])
local windowSize = tonumber(ARGV[2])
local maxRequests = tonumber(ARGV[3])
local requestID = ARGV[4]

local windowStart = currentTime - windowSize + 1

redis.call(
    "ZREMRANGEBYSCORE",
    windowKey,
    "-inf",
    windowStart
)

local requestCount = redis.call(
    "ZCARD",
    windowKey
)

if requestCount < maxRequests then

    redis.call(
        "ZADD",
        windowKey,
        currentTime,
        requestID
    )

    return {
        1,
        maxRequests - requestCount - 1,
        0
    }

else

    local oldestRequest = redis.call(
        "ZRANGE",
        windowKey,
        0,
        0,
        "WITHSCORES"
    )

    local oldestTimestamp = tonumber(oldestRequest[2])

    local retryAfter = oldestTimestamp + windowSize - currentTime

    if retryAfter < 0 then
        retryAfter = 0
    end

    return {
        0,
        0,
        retryAfter
    }

end
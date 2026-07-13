package source

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRateLimiter(
	config ClientConfig,
	client *redis.Client,
	tokenBucketScript string,
	slidingWindowScript string,
	windowSize int64,
) (RateLimiter, error) {

	switch config.Algorithm {

	case "token_bucket":
		return &TokenBucketLimiter{
			Client: client,
			Script: tokenBucketScript,
		}, nil

	case "sliding_window":
		return &SlidingWindowLimiter{
			Client:     client,
			Script:     slidingWindowScript,
			WindowSize: windowSize,
		}, nil

	default:
		return nil, fmt.Errorf("unknown rate limiting algorithm")
	}
}

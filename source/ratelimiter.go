package source

import "context"

type RateLimiter interface {
	Allow(
		ctx context.Context,
		config ClientConfig,
	) (interface{}, error)
}

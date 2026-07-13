package source

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenBucketLimiter struct {
	Client *redis.Client
	Script string
}

func (t *TokenBucketLimiter) Allow(
	ctx context.Context,
	config ClientConfig,
) (interface{}, error) {

	bucketKey := "bucket:" + config.ClientID

	return t.Client.Eval(
		ctx,
		t.Script,
		[]string{
			bucketKey,
		},
		time.Now().Unix(),
		config.Burst,
		config.Rate,
	).Result()
}

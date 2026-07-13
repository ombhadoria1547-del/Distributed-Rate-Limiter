package source

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type SlidingWindowLimiter struct {
	Client     *redis.Client
	Script     string
	WindowSize int64
}

func (s *SlidingWindowLimiter) Allow(
	ctx context.Context,
	config ClientConfig,
) (interface{}, error) {

	windowKey := "rl:window:" + config.ClientID

	return s.Client.Eval(
		ctx,
		s.Script,
		[]string{
			windowKey,
		},
		time.Now().Unix(),
		s.WindowSize,
		config.Burst,
		config.ClientID+"-"+strconv.FormatInt(time.Now().UnixNano(), 10),
	).Result()
}

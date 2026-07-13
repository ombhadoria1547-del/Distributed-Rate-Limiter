package source

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func SaveClientConfig(
	ctx context.Context,
	client *redis.Client,
	config ClientConfig,
) error {

	key := "rl:cfg:" + config.ClientID

	return client.HSet(
		ctx,
		key,
		map[string]interface{}{
			"rate":      strconv.FormatFloat(config.Rate, 'f', -1, 64),
			"burst":     strconv.FormatFloat(config.Burst, 'f', -1, 64),
			"algorithm": config.Algorithm,
		},
	).Err()
}

func GetClientConfig(
	ctx context.Context,
	client *redis.Client,
	clientID string,
) (ClientConfig, error) {

	key := "rl:cfg:" + clientID

	result, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		return ClientConfig{}, err
	}

	if len(result) == 0 {
		return ClientConfig{}, redis.Nil
	}

	rate, err := strconv.ParseFloat(result["rate"], 64)
	if err != nil {
		return ClientConfig{}, err
	}

	burst, err := strconv.ParseFloat(result["burst"], 64)
	if err != nil {
		return ClientConfig{}, err
	}

	algorithm := result["algorithm"]

	if algorithm == "" {
		algorithm = "token_bucket"
	}

	return ClientConfig{
		ClientID:  clientID,
		Rate:      rate,
		Burst:     burst,
		Algorithm: algorithm,
	}, nil
}

func GetAllClientConfigs(
	ctx context.Context,
	client *redis.Client,
) ([]ClientConfig, error) {

	var configs []ClientConfig
	var cursor uint64

	for {

		keys, nextCursor, err := client.Scan(ctx, cursor, "rl:cfg:*", 10).Result()
		if err != nil {
			return nil, err
		}

		for _, key := range keys {

			clientID := key[len("rl:cfg:"):]

			config, err := GetClientConfig(ctx, client, clientID)
			if err != nil {
				continue
			}

			configs = append(configs, config)
		}

		cursor = nextCursor

		if cursor == 0 {
			break
		}
	}

	return configs, nil
}

func UpdateClientConfig(
	ctx context.Context,
	client *redis.Client,
	config ClientConfig,
) error {

	key := "rl:cfg:" + config.ClientID

	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		return err
	}

	if exists == 0 {
		return redis.Nil
	}

	return client.HSet(
		ctx,
		key,
		map[string]interface{}{
			"rate":      strconv.FormatFloat(config.Rate, 'f', -1, 64),
			"burst":     strconv.FormatFloat(config.Burst, 'f', -1, 64),
			"algorithm": config.Algorithm,
		},
	).Err()
}

func DeleteClientConfig(
	ctx context.Context,
	client *redis.Client,
	clientID string,
) error {

	key := "rl:cfg:" + clientID

	deleted, err := client.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	if deleted == 0 {
		return redis.Nil
	}

	return nil
}

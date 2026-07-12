package main

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ombhadoria1547-del/Distributed-Rate-Limiter/source"
	"github.com/redis/go-redis/v9"
)

type CheckResponse struct {
	Allowed    bool  `json:"allowed"`
	Remaining  int64 `json:"remaining"`
	RetryAfter int64 `json:"retry_after"`
}

func main() {

	router := gin.Default()

	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	ctx := context.Background()

	err := client.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	script, err := os.ReadFile("scripts/token_bucket.lua")
	if err != nil {
		panic(err)
	}

	defaultCapacity := 10.0
	defaultRefillRate := 1.0

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	router.GET("/user", func(c *gin.Context) {

		name := c.Query("name")

		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "name query parameter is required",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": name,
		})
	})

	router.GET("/user/:id", func(c *gin.Context) {

		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	router.POST("/admin/clients", func(c *gin.Context) {

		var config source.ClientConfig

		if err := c.ShouldBindJSON(&config); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body",
			})
			return
		}

		if config.ClientID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "client_id is required",
			})
			return
		}

		err := source.SaveClientConfig(ctx, client, config)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to save client configuration",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "client created",
			"client":  config,
		})
	})

	router.GET("/admin/clients", func(c *gin.Context) {

		configs, err := source.GetAllClientConfigs(ctx, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to fetch client configurations",
			})
			return
		}

		c.JSON(http.StatusOK, configs)

	})

	router.PUT("/admin/clients", func(c *gin.Context) {

		var config source.ClientConfig

		if err := c.ShouldBindJSON(&config); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body",
			})
			return
		}

		if config.ClientID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "client_id is required",
			})
			return
		}

		err := source.UpdateClientConfig(ctx, client, config)

		if err == redis.Nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "client not found",
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update client configuration",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "client updated",
			"client":  config,
		})
	})

	router.DELETE("/admin/clients", func(c *gin.Context) {

		clientID := c.Query("client_id")

		if clientID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "client_id query parameter is required",
			})
			return
		}

		err := source.DeleteClientConfig(ctx, client, clientID)

		if err == redis.Nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "client not found",
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to delete client configuration",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "client deleted",
		})
	})

	router.GET("/check", func(c *gin.Context) {

		clientID := c.Query("client_id")

		if clientID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "client_id query parameter is required",
			})
			return
		}

		capacity := defaultCapacity
		refillRate := defaultRefillRate

		config, err := source.GetClientConfig(ctx, client, clientID)
		if err != nil && err != redis.Nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to load client configuration",
			})
			return
		}

		if err == nil {
			capacity = config.Burst
			refillRate = config.Rate
		}

		bucketKey := "bucket:" + clientID

		result, err := client.Eval(
			ctx,
			string(script),
			[]string{
				bucketKey,
			},
			time.Now().Unix(),
			capacity,
			refillRate,
		).Result()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to execute rate limiter",
			})
			return
		}

		response, ok := result.([]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "unexpected response from Lua script",
			})
			return
		}

		if len(response) != 3 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "unexpected response from Lua script",
			})
			return
		}

		allowed, ok := response[0].(int64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "unexpected response from Lua script",
			})
			return
		}

		remaining, ok := response[1].(int64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "unexpected response from Lua script",
			})
			return
		}

		retryAfter, ok := response[2].(int64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "unexpected response from Lua script",
			})
			return
		}

		allowedBool := allowed == 1

		c.Header(
			"X-RateLimit-Limit",
			strconv.FormatFloat(capacity, 'f', 0, 64),
		)

		c.Header(
			"X-RateLimit-Remaining",
			strconv.FormatInt(remaining, 10),
		)

		c.Header(
			"Retry-After",
			strconv.FormatInt(retryAfter, 10),
		)

		if allowedBool {

			c.JSON(http.StatusOK, CheckResponse{
				Allowed:    allowedBool,
				Remaining:  remaining,
				RetryAfter: retryAfter,
			})

		} else {

			c.JSON(http.StatusTooManyRequests, CheckResponse{
				Allowed:    allowedBool,
				Remaining:  remaining,
				RetryAfter: retryAfter,
			})

		}

	})

	router.Run(":8080")
}

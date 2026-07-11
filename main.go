package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {

	router := gin.Default()

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
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

	capacity := 10.0
	refillRate := 1.0

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

	router.GET("/check", func(c *gin.Context) {

		clientID := c.Query("client_id")

		if clientID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "client_id query parameter is required",
			})
			return
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

		allowed, ok := result.(int64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "unexpected response from Lua script",
			})
			return
		}

		if allowed == 1 {

			c.JSON(http.StatusOK, gin.H{
				"allowed": true,
			})

		} else {

			c.JSON(http.StatusTooManyRequests, gin.H{
				"allowed": false,
			})

		}

	})
	router.Run(":8080")
}

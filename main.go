package main

import (
	"context"
	"fmt"
	"net/http"

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

	err = client.Set(ctx, "username", "om", 0).Err()
	if err != nil {
		panic(err)
	}

	value, err := client.Get(ctx, "username").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(value)

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

	router.Run(":8080")
}

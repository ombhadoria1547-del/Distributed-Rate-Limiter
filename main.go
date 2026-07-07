package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

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

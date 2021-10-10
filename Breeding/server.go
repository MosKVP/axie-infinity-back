package main

import (
	"breeding/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/calculate", api.Calculate)
	r.Run() // listen and serve on localhost:8080
}

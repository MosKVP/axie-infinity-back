package main

import (
	"axie-infinity-back/internal/api/calculate"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	state := os.Getenv("STATE")

	var envFilename string
	switch state {
	case "PROD":
		envFilename = "prod.env"
		gin.SetMode(gin.ReleaseMode)
	default:
		envFilename = "local.env"
	}
	err := godotenv.Load(envFilename)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Accept", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Path
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/v1/breeding/calculate", calculate.Calculate)

	log.Println("Server Started")
	r.Run() // listen and serve on localhost:8080
}

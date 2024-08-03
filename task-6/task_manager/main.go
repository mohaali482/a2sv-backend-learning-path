package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/router"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getHost() string {
	hostUrl := os.Getenv("HOST_URL")
	if hostUrl != "" {
		return hostUrl
	}

	return "localhost:8000"
}

func main() {
	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	s := data.NewMongoTaskService()
	router.RegisterHandlers(r, s)

	r.Run(getHost())
}

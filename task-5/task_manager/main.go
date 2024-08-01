package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohaali482/a2sv-backend-learning-path/task-5/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-5/router"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	s := data.NewMongoTaskService()
	router.RegisterHandlers(r, s)

	r.Run("localhost:8000")
}

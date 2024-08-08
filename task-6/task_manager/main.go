package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/data/mongodb"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/router"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/validator"
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

	c := mongodb.GetNewMongoClient()
	ts := mongodb.NewMongoTaskService(c)
	us := mongodb.NewMongoUserService(c)

	r := gin.New()
	validator.BindingValidator()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	router.RegisterTaskHandlers(r, ts, us)
	router.RegisterUserHandlers(r, us)

	r.Run(getHost())
}

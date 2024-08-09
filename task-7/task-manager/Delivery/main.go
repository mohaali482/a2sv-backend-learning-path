package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Delivery/routers"
	"github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Domain/validator"
	"github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Repositories/mongodb"
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
	r := gin.New()
	validator.BindingValidator()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routers.Setup(c.Database(os.Getenv("MONGO_DB")), r)

	r.Run(getHost())
}

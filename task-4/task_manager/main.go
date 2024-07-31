package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mohaali482/a2sv-backend-learning-path/task-4/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-4/router"
)

func main() {
	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	s := data.NewTaskService()
	router.RegisterHandlers(r, s)

	r.Run("localhost:8000")
}

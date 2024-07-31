package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohaali482/a2sv-backend-learning-path/task-4/data"
)

func GetAllTasksController(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, s.GetAllTasks())
	}
}

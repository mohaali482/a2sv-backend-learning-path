package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohaali482/a2sv-backend-learning-path/task-4/data"
)

func GetAllTasksController(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, s.GetAllTasks())
	}
}

func GetTaskById(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id",
			})
			return
		}

		task, err := s.GetTaskById(intId)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "task not found",
			})
		}

		c.JSON(http.StatusOK, task)
	}
}

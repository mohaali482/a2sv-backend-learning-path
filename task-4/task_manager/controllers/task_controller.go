package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohaali482/a2sv-backend-learning-path/task-4/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-4/models"
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

func UpdateTask(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id",
			})
			return
		}

		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		updatedTask, err := s.UpdateTask(intId, task)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "task not found",
			})
		}

		c.JSON(http.StatusOK, updatedTask)
	}
}

func DeleteTask(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id",
			})
			return
		}

		err = s.DeleteTask(intId)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "task not found",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "task deleted",
		})
	}
}

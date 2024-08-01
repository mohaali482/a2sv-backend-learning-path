package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohaali482/a2sv-backend-learning-path/task-5/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-5/models"
)

func GetAllTasks(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, s.GetAllTasks(c.Request.Context()))
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

		task, err := s.GetTaskById(c.Request.Context(), intId)

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

		task, err = s.UpdateTask(c.Request.Context(), intId, task)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "task not found",
			})
		}

		c.JSON(http.StatusOK, task)
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

		err = s.DeleteTask(c.Request.Context(), intId)

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

func CreateTask(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		task = s.CreateTask(c.Request.Context(), task)

		c.JSON(http.StatusCreated, task)
	}
}

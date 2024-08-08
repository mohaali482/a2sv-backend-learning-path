package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ve "github.com/go-playground/validator/v10"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/models"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskForm struct {
	UserId      string    `json:"user_id" binding:"required,mongodb"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	Done        bool      `json:"done"`
}

func (t *taskForm) ToModelTask() (models.Task, error) {
	userObjectId, err := primitive.ObjectIDFromHex(t.UserId)
	if err != nil {
		return models.Task{}, errors.New("invalid user id")
	}

	return models.Task{
		UserId:      userObjectId,
		Title:       t.Title,
		Description: t.Description,
		DateTime:    t.DateTime,
		Done:        t.Done,
	}, nil
}

func GetAllTasks(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authenticatedUser, _ := c.Get("user")
		user := authenticatedUser.(*models.User)

		if user.GetRole() == "admin" {
			c.JSON(http.StatusOK, s.GetAllTasks(c.Request.Context()))
		} else {
			c.JSON(http.StatusOK, s.GetUserTasks(c.Request.Context(), user.ID.Hex()))
		}

	}
}

func GetTaskById(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		authenticatedUser, _ := c.Get("user")
		user := authenticatedUser.(*models.User)

		var task *models.Task
		var err error
		if user.GetRole() == "admin" {
			task, err = s.GetTaskById(c.Request.Context(), id)
		} else {
			task, err = s.GetUserTaskById(c.Request.Context(), id, user.ID.Hex())
		}

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "task not found",
			})
			return
		}

		c.JSON(http.StatusOK, task)
	}
}

func UpdateTask(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var updatedTask taskForm

		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			var ver ve.ValidationErrors
			if errors.As(err, &ver) {
				validator.ReturnErrorResponse(err, c)
			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
			return
		}

		task, err := updatedTask.ToModelTask()

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task, err = s.UpdateTask(c.Request.Context(), id, task)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "task not found",
			})
			return
		}

		c.JSON(http.StatusOK, task)
	}
}

func DeleteTask(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := s.DeleteTask(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "task not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "task deleted",
		})
	}
}

func CreateTask(s data.TaskUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createTask taskForm

		if err := c.ShouldBindJSON(&createTask); err != nil {
			var ver ve.ValidationErrors
			if errors.As(err, &ver) {
				validator.ReturnErrorResponse(err, c)
			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
			return
		}

		task, err := createTask.ToModelTask()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task, err = s.CreateTask(c.Request.Context(), task)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, task)
	}
}

package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ve "github.com/go-playground/validator/v10"
	domain "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Domain"
	"github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Domain/validator"
	usecases "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Usecases"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskForm struct {
	UserId      string    `json:"user_id" binding:"required,mongodb"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	Done        bool      `json:"done"`
}

func (t *taskForm) ToModelTask() (domain.Task, error) {
	userObjectId, err := primitive.ObjectIDFromHex(t.UserId)
	if err != nil {
		return domain.Task{}, errors.New("invalid user id")
	}

	return domain.Task{
		UserId:      userObjectId,
		Title:       t.Title,
		Description: t.Description,
		DateTime:    t.DateTime,
		Done:        t.Done,
	}, nil
}

type TaskController struct {
	TaskUsecase usecases.TaskUsecase
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	authenticatedUser, _ := c.Get("user")
	user := authenticatedUser.(*domain.User)

	if user.GetRole() == "admin" {
		c.JSON(http.StatusOK, tc.TaskUsecase.GetAllTasks(c.Request.Context()))
	} else {
		c.JSON(http.StatusOK, tc.TaskUsecase.GetUserTasks(c.Request.Context(), user.ID.Hex()))
	}

}

func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	authenticatedUser, _ := c.Get("user")
	user := authenticatedUser.(*domain.User)

	var task *domain.Task
	var err error
	if user.GetRole() == "admin" {
		task, err = tc.TaskUsecase.GetTaskById(c.Request.Context(), id)
	} else {
		task, err = tc.TaskUsecase.GetUserTaskById(c.Request.Context(), id, user.ID.Hex())
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
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

	task, err = tc.TaskUsecase.UpdateTask(c.Request.Context(), id, task)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := tc.TaskUsecase.DeleteTask(c.Request.Context(), id)

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

func (tc *TaskController) CreateTask(c *gin.Context) {
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

	task, err = tc.TaskUsecase.CreateTask(c.Request.Context(), task)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

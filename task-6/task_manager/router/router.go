package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/controllers"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/data"
)

func RegisterHandlers(r *gin.Engine, s data.TaskUseCase) {
	r.GET("/tasks", controllers.GetAllTasks(s))
	r.GET("/tasks/:id", controllers.GetTaskById(s))
	r.PUT("/tasks/:id", controllers.UpdateTask(s))
	r.DELETE("/tasks/:id", controllers.DeleteTask(s))
	r.POST("/tasks", controllers.CreateTask(s))
}

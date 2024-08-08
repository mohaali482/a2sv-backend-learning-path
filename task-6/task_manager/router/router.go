package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/controllers"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/middleware"
)

func RegisterTaskHandlers(r *gin.Engine, ts data.TaskUseCase, us data.UserUseCase) {
	taskRouter := r.Group("/tasks", middleware.AuthMiddleware(us))

	{
		taskRouter.GET("", controllers.GetAllTasks(ts))
		taskRouter.GET("/:id", controllers.GetTaskById(ts))

		taskRouter.Use(middleware.AdminMiddleware(us))
		taskRouter.PUT("/:id", controllers.UpdateTask(ts))
		taskRouter.DELETE("/:id", controllers.DeleteTask(ts))
		taskRouter.POST("", controllers.CreateTask(ts))
	}
}

func RegisterUserHandlers(r *gin.Engine, us data.UserUseCase) {
	r.POST("/login", controllers.Login(us))
	r.POST("/register", controllers.Register(us))

	r.Use(middleware.AdminMiddleware(us))
	r.POST("/promote", controllers.Promote(us))
}

package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Delivery/controllers"
	"github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Infrastructure/middlewares"
	repositories "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Repositories"
	usecases "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskRouter(db *mongo.Database, group *gin.RouterGroup) {
	tr := repositories.NewMongoTaskRepository(db, os.Getenv("MONGO_TASK_COLLECTION"))
	ur := repositories.NewMongoUserRepository(db, os.Getenv("MONGO_USER_COLLECTION"))
	us := usecases.UserUsecase{
		Repository: ur,
		Key:        []byte(os.Getenv("JWT_KEY")),
	}

	tc := controllers.TaskController{
		TaskUsecase: usecases.TaskUsecase{
			Repository:  tr,
			UserUsecase: us,
		},
	}

	group.Use(middlewares.AuthMiddleware(us))
	{
		group.GET("", tc.GetAllTasks)
		group.GET("/:id", tc.GetTaskById)

		group.Use(middlewares.AdminMiddleware(us))
		group.PUT("/:id", tc.UpdateTask)
		group.DELETE("/:id", tc.DeleteTask)
		group.POST("", tc.CreateTask)
	}
}

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

func NewUserRouter(db *mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewMongoUserRepository(db, os.Getenv("MONGO_USER_COLLECTION"))
	us := usecases.UserUsecase{
		Repository: ur,
		Key:        []byte(os.Getenv("JWT_KEY")),
	}

	tc := controllers.UserControllers{
		UserUsecase: us,
	}

	group.POST("/login", tc.Login)
	group.POST("/register", tc.Register)

	group.Use(middlewares.AuthMiddleware(us))
	group.Use(middlewares.AdminMiddleware(us))
	group.POST("/promote", tc.Promote)
}

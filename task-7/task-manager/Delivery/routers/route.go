package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db *mongo.Database, gin *gin.Engine) {
	publicGroup := gin.Group("")
	NewUserRouter(db, publicGroup)

	taskGroup := gin.Group("tasks")
	NewTaskRouter(db, taskGroup)
}

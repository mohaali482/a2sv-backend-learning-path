package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Domain"
	usecases "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Usecases"
)

func AdminMiddleware(s usecases.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		savedUser, exists := c.Get("user")
		user := savedUser.(*domain.User)

		if !exists || user.GetRole() != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "you're not authorized",
			})
			return
		}

		c.Next()
	}
}

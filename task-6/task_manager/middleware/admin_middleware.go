package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/models"
)

func AdminMiddleware(s data.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		savedUser, exists := c.Get("user")
		user := savedUser.(*models.User)

		if !exists || user.GetRole() != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "you're not authorized",
			})
			return
		}

		c.Next()
	}
}

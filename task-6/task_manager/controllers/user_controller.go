package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/validator"
)

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PromoteForm struct {
	Username string `json:"username" binding:"required"`
}

func Login(s data.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginForm LoginForm

		err := c.ShouldBind(&loginForm)
		if err != nil {
			validator.ReturnErrorResponse(err, c)
			return
		}

		token, err := s.Login(c.Request.Context(), loginForm.Username, loginForm.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}

func Register(s data.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerForm RegisterForm

		err := c.ShouldBind(&registerForm)
		if err != nil {
			validator.ReturnErrorResponse(err, c)
			return
		}

		token, err := s.Register(c.Request.Context(), registerForm.Username, registerForm.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}

func Promote(s data.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var promoteForm PromoteForm

		err := c.ShouldBind(&promoteForm)
		if err != nil {
			validator.ReturnErrorResponse(err, c)
			return
		}

		err = s.Promote(c.Request.Context(), promoteForm.Username)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "user promoted succesfuly"})
	}
}

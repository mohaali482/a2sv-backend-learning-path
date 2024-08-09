package infrastructure

import (
	"github.com/golang-jwt/jwt"
	domain "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Domain"
)

func NewJWTSignedString(key []byte, user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID.Hex(),
		"username": user.Username,
	})

	return token.SignedString(key)
}

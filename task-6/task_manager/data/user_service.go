package data

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/models"
)

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyPromoted = errors.New("user is already promoted")
var ErrUniqueUsername = errors.New("user with this username already exists")

type UserUseCase interface {
	Login(ctx context.Context, username string, password string) (string, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	VerifyToken(ctx context.Context, tokenString string) (jwt.MapClaims, error)
	Promote(ctx context.Context, username string) error
	Register(ctx context.Context, username string, password string) (string, error)
}

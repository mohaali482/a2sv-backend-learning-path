package usecases

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	domain "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Domain"
	infrastructure "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Infrastructure"
	repositories "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	Repository repositories.UserRepository
	Key        []byte
}

func (s *UserUsecase) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := s.Repository.GetUserByUsername(ctx, username)

	if err != nil {
		return "", err
	}

	valid := infrastructure.CompareHashAndPassword(user.Password, password)

	if !valid {
		return "", domain.ErrInvalidCredentials
	}

	tokenString, err := infrastructure.NewJWTSignedString(s.Key, *user)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (s *UserUsecase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.Repository.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserUsecase) VerifyToken(ctx context.Context, tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.Key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func (s *UserUsecase) Promote(ctx context.Context, username string) error {
	user, err := s.Repository.GetUserByUsername(ctx, username)

	if err != nil {
		return err
	}

	if user.GetRole() == "admin" {
		return domain.ErrUserAlreadyPromoted
	}

	user.Role = 1
	return s.Repository.UpdateUser(ctx, user.ID.Hex(), *user)
}

func (s *UserUsecase) Register(ctx context.Context, username string, password string) (string, error) {
	count, err := s.Repository.GetUsersCount(ctx)
	if err != nil {
		log.Fatal(err)
	}

	user, err := s.Repository.GetUserByUsername(ctx, username)

	if err != nil {
		return "", domain.ErrUniqueUsername
	}

	if count == 0 {
		user.Role = 0

	} else {
		user.Role = 1
	}

	user.Username = strings.ToLower(user.Username)
	userID, err := s.Repository.CreateUser(ctx, *user)
	if err != nil {
		return "", err
	}

	userObjectID, _ := primitive.ObjectIDFromHex(userID)
	user.ID = userObjectID

	return s.Login(ctx, user.Username, password)
}

func (s *UserUsecase) IsValidId(ctx context.Context, id string) error {
	_, err := s.Repository.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

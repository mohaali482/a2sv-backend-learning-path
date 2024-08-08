package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type MongoUserService struct {
	client *mongo.Client
	key    []byte
}

func NewMongoUserService(client *mongo.Client) *MongoUserService {
	key := os.Getenv("JWT_KEY")
	if key == "" {
		log.Fatal("JWT_KEY not set")
	}
	return &MongoUserService{client: client, key: []byte(key)}
}

func (m *MongoUserService) Login(ctx context.Context, username string, password string) (string, error) {
	collection := m.client.Database("taskManager").Collection("users")
	var user models.User

	err := collection.FindOne(ctx, bson.D{{Key: "username", Value: strings.ToLower(username)}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", data.ErrInvalidCredentials
		}

		log.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return "", data.ErrInvalidCredentials
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID.Hex(),
		"username": username,
	})

	tokenString, err := token.SignedString(m.key)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (s *MongoUserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	collection := s.client.Database("taskManager").Collection("users")
	userObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, data.ErrInvalidUserId
	}

	var user models.User

	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: userObjectId}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.ErrUserNotFound
		} else {
			log.Fatal(err)
		}
	}

	return &user, nil
}

func (m *MongoUserService) VerifyToken(ctx context.Context, tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return m.key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func (m *MongoUserService) Promote(ctx context.Context, username string) error {
	collection := m.client.Database("taskManager").Collection("users")
	var user models.User

	err := collection.FindOne(ctx, bson.D{{Key: "username", Value: strings.ToLower(username)}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return data.ErrUserNotFound
		}

		log.Fatal(err)
	}

	if user.GetRole() == "admin" {
		return data.ErrUserAlreadyPromoted
	}

	_, err = collection.UpdateByID(ctx, user.ID, bson.D{{
		Key: "$set", Value: bson.D{{
			Key: "role", Value: 0,
		}}}})

	if err != nil {
		return err
	}

	return nil

}

func (m *MongoUserService) Register(ctx context.Context, username string, password string) (string, error) {
	collections := m.client.Database("taskManager").Collection("users")
	count, err := collections.EstimatedDocumentCount(ctx)
	if err != nil {
		log.Fatal(err)
	}

	result := collections.FindOne(ctx, bson.D{{Key: "username", Value: strings.ToLower(username)}})

	if result.Err() != mongo.ErrNoDocuments {
		return "", data.ErrUniqueUsername
	}

	var user *models.User
	if count == 0 {
		user, err = m.registerUser(ctx, username, password, 0)
	} else {
		user, err = m.registerUser(ctx, username, password, 1)
	}

	if err != nil {
		return "", err
	}

	return m.Login(ctx, user.Username, password)
}

func (m *MongoUserService) registerUser(ctx context.Context, username string, password string, admin int) (*models.User, error) {
	collections := m.client.Database("taskManager").Collection("users")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	result, err := collections.InsertOne(ctx, models.User{
		Username: strings.ToLower(username),
		Password: string(hashedPassword),
		Role:     admin,
	})

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	err = collections.FindOne(ctx, bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.ErrUserNotFound
		}

		log.Fatal(err)
	}

	return &user, nil

}

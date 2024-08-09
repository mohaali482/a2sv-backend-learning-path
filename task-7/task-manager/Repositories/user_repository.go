package repositories

import (
	"context"

	domain "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (string, error)
	UpdateUser(ctx context.Context, id string, user domain.User) error
}

type MongoUserRepository struct {
	db         *mongo.Database
	collection string
}

func NewMongoUserRepository(db *mongo.Database, collection string) *MongoUserRepository {
	return &MongoUserRepository{
		db:         db,
		collection: collection,
	}
}

func (s *MongoUserRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	collection := s.db.Collection(s.collection)
	userObjectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, domain.ErrInvalidUserId
	}

	var user domain.User
	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: userObjectId}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (s *MongoUserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	collection := s.db.Collection(s.collection)

	var user domain.User
	err := collection.FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (s *MongoUserRepository) CreateUser(ctx context.Context, user domain.User) (string, error) {
	collection := s.db.Collection(s.collection)
	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *MongoUserRepository) UpdateUser(ctx context.Context, id string, user domain.User) error {
	collection := s.db.Collection(s.collection)
	userObjectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return domain.ErrInvalidUserId
	}

	updateResult, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: userObjectId}}, user)
	if updateResult.MatchedCount == 0 {
		return domain.ErrUserNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

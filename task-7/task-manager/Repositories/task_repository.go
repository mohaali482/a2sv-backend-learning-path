package repositories

import (
	"context"
	"log"

	domain "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	GetAllTasks(ctx context.Context) []*domain.Task
	GetUserTasks(ctx context.Context, userId string) []*domain.Task
	GetUserTaskById(ctx context.Context, id string, userId string) (*domain.Task, error)
	GetTaskById(ctx context.Context, id string) (*domain.Task, error)
	CreateTask(ctx context.Context, task domain.Task) (string, error)
	UpdateTask(ctx context.Context, id string, task domain.Task) error
	DeleteTask(ctx context.Context, id string) error
}

type MongoTaskRepository struct {
	db         *mongo.Database
	collection string
}

func NewMongoTaskRepository(db *mongo.Database, collection string) *MongoTaskRepository {
	return &MongoTaskRepository{
		db:         db,
		collection: collection,
	}
}

func (s *MongoTaskRepository) GetAllTasks(ctx context.Context) []*domain.Task {
	collection := s.db.Collection(s.collection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var tasks []*domain.Task = make([]*domain.Task, 0)

	if err = cursor.All(ctx, &tasks); err != nil {
		log.Fatal(err)
	}

	return tasks
}

func (s *MongoTaskRepository) GetUserTasks(ctx context.Context, userId string) []*domain.Task {
	collection := s.db.Collection(s.collection)
	userObjectId, _ := primitive.ObjectIDFromHex(userId)

	cursor, err := collection.Find(ctx, bson.D{{Key: "user_id", Value: userObjectId}})
	if err != nil {
		log.Fatal(err)
	}

	var tasks []*domain.Task = make([]*domain.Task, 0)
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		log.Fatal(err)
	}

	return tasks
}

func (s *MongoTaskRepository) GetUserTaskById(ctx context.Context, id string, userId string) (*domain.Task, error) {
	collection := s.db.Collection(s.collection)

	taskObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrInvalidTaskId
	}

	userObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrInvalidUserId
	}

	var task domain.Task

	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: taskObjectId}, {Key: "user_id", Value: userObjectId}}).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrTaskNotFound
		} else {
			return nil, err
		}
	}

	return &task, nil
}

func (s *MongoTaskRepository) GetTaskById(ctx context.Context, id string) (*domain.Task, error) {
	collection := s.db.Collection(s.collection)

	taskObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrInvalidTaskId
	}

	var task domain.Task

	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: taskObjectId}}).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrTaskNotFound
		} else {
			return nil, err
		}
	}

	return &task, nil
}

func (s *MongoTaskRepository) UpdateTask(ctx context.Context, id string, task domain.Task) error {
	taskObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrInvalidTaskId
	}

	collection := s.db.Collection(s.collection)

	updateResult, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: taskObjectId}}, task)

	if updateResult.MatchedCount == 0 {
		return domain.ErrTaskNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *MongoTaskRepository) DeleteTask(ctx context.Context, id string) error {
	taskObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrInvalidTaskId
	}

	collection := s.db.Collection(s.collection)

	deleteResult, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: taskObjectId}})

	if deleteResult.DeletedCount == 0 {
		return domain.ErrTaskNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *MongoTaskRepository) CreateTask(ctx context.Context, task domain.Task) (string, error) {
	collection := s.db.Collection(s.collection)

	result, err := collection.InsertOne(ctx, task)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

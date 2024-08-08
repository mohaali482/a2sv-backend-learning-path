package mongodb

import (
	"context"
	"errors"
	"log"

	"github.com/mohaali482/a2sv-backend-learning-path/task-6/data"
	"github.com/mohaali482/a2sv-backend-learning-path/task-6/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskService struct {
	client *mongo.Client
}

func NewMongoTaskService(client *mongo.Client) *MongoTaskService {
	return &MongoTaskService{client: client}
}

func (s *MongoTaskService) GetAllTasks(ctx context.Context) []*models.Task {
	collection := s.client.Database("taskManager").Collection("tasks")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var tasks []*models.Task = make([]*models.Task, 0)

	if err = cursor.All(ctx, &tasks); err != nil {
		log.Fatal(err)
	}

	return tasks
}

func (s *MongoTaskService) GetUserTasks(ctx context.Context, userId string) []*models.Task {
	collection := s.client.Database("taskManager").Collection("tasks")
	userObjectId, _ := primitive.ObjectIDFromHex(userId)

	cursor, err := collection.Find(ctx, bson.D{{Key: "user_id", Value: userObjectId}})
	if err != nil {
		log.Fatal(err)
	}

	var tasks []*models.Task = make([]*models.Task, 0)
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		log.Fatal(err)
	}

	return tasks
}

func (s *MongoTaskService) GetUserTaskById(ctx context.Context, id string, userId string) (*models.Task, error) {
	collection := s.client.Database("taskManager").Collection("tasks")
	taskObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, data.ErrInvalidTaskId
	}
	userObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, data.ErrInvalidUserId
	}

	var task models.Task

	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: taskObjectId}, {Key: "user_id", Value: userObjectId}}).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		} else {
			log.Fatal(err)
		}
	}

	return &task, nil
}

func (s *MongoTaskService) GetTaskById(ctx context.Context, id string) (*models.Task, error) {
	collection := s.client.Database("taskManager").Collection("tasks")
	taskObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, data.ErrInvalidTaskId
	}

	var task models.Task

	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: taskObjectId}}).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		} else {
			log.Fatal(err)
		}
	}

	return &task, nil
}

func (s *MongoTaskService) UpdateTask(ctx context.Context, id string, task models.Task) (models.Task, error) {
	if _, err := s.GetTaskById(ctx, id); err != nil {
		return models.Task{}, errors.New("task not found")
	}

	taskObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, data.ErrInvalidTaskId
	}

	collection := s.client.Database("taskManager").Collection("tasks")

	_, err = collection.UpdateOne(ctx, bson.D{{Key: "_id", Value: taskObjectId}}, bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "title", Value: task.Title},
			{Key: "done", Value: task.Done},
			{Key: "description", Value: task.Description},
			{Key: "datetime", Value: task.DateTime},
			{Key: "user_id", Value: task.UserId},
		},
	}})

	if err != nil {
		log.Fatal(err)
	}

	updatedTask, err := s.GetTaskById(ctx, id)

	return *updatedTask, err
}

func (s *MongoTaskService) DeleteTask(ctx context.Context, id string) error {
	taskObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return data.ErrInvalidTaskId
	}

	collection := s.client.Database("taskManager").Collection("tasks")

	result, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: taskObjectId}})

	if err != nil {
		log.Fatal(err)
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (s *MongoTaskService) CreateTask(ctx context.Context, task models.Task) models.Task {
	collection := s.client.Database("taskManager").Collection("tasks")

	result, err := collection.InsertOne(ctx, task)

	if err != nil {
		log.Fatal(err)
	}

	var insertedTask models.Task

	sResult := collection.FindOne(ctx, bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&insertedTask)

	if sResult != nil {
		log.Fatal(sResult)
	}

	return insertedTask
}

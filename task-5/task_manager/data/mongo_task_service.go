package data

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/mohaali482/a2sv-backend-learning-path/task-5/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTaskService struct {
	client *mongo.Client
}

func NewMongoTaskService() *MongoTaskService {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoTaskService{client: client}
}

func (s *MongoTaskService) GetAllTasks(ctx context.Context) []*models.Task {
	collection := s.client.Database("taskManager").Collection("tasks")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var tasks []*models.Task
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var elem models.Task
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, &elem)
	}

	return tasks
}

func (s *MongoTaskService) GetTaskById(ctx context.Context, id int) (*models.Task, error) {
	collection := s.client.Database("taskManager").Collection("tasks")

	var task models.Task

	err := collection.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		} else {
			log.Fatal(err)
		}
	}

	return &task, nil
}

func (s *MongoTaskService) UpdateTask(ctx context.Context, id int, task models.Task) (models.Task, error) {
	if _, err := s.GetTaskById(ctx, id); err != nil {
		return models.Task{}, errors.New("task not found")
	}

	collection := s.client.Database("taskManager").Collection("tasks")

	_, err := collection.UpdateOne(ctx, bson.D{{Key: "id", Value: id}}, bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "title", Value: task.Title},
			{Key: "done", Value: task.Done},
		},
	}})

	if err != nil {
		log.Fatal(err)
	}

	updatedTask, err := s.GetTaskById(ctx, id)

	return *updatedTask, err
}

func (s *MongoTaskService) DeleteTask(ctx context.Context, id int) error {
	collection := s.client.Database("taskManager").Collection("tasks")

	result, err := collection.DeleteOne(ctx, bson.D{{Key: "id", Value: id}})

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
	counterCollection := s.client.Database("taskManager").Collection("counters")
	var counter struct {
		SeqValue int `bson:"seq_value"`
	}

	err := counterCollection.FindOneAndUpdate(ctx, bson.D{}, bson.D{{
		Key: "$inc",
		Value: bson.D{{
			Key:   "seq_value",
			Value: 1,
		}},
	}}).Decode(&counter)

	if err != nil {
		log.Fatal(err)
	}

	task.Id = counter.SeqValue

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

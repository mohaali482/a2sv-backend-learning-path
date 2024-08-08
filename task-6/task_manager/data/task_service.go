package data

import (
	"context"
	"errors"

	"github.com/mohaali482/a2sv-backend-learning-path/task-6/models"
)

var ErrInvalidUserId = errors.New("invalid user id")
var ErrInvalidTaskId = errors.New("invalid task id")

type TaskUseCase interface {
	GetAllTasks(ctx context.Context) []*models.Task
	GetUserTasks(ctx context.Context, userId string) []*models.Task
	GetTaskById(ctx context.Context, id string) (*models.Task, error)
	GetUserTaskById(ctx context.Context, id string, userId string) (*models.Task, error)
	UpdateTask(ctx context.Context, id string, task models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, id string) error
	CreateTask(ctx context.Context, task models.Task) (models.Task, error)
}

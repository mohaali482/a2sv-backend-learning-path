package repository

import (
	"context"

	"github.com/mohaali482/a2sv-backend-learning-path/task-6/models"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	GetTask(ctx context.Context, id int) (*models.Task, error)
	GetTasks(ctx context.Context) ([]*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	DeleteTask(ctx context.Context, id int) error
}

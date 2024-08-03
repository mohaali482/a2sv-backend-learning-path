package repository

import "github.com/mohaali482/a2sv-backend-learning-path/task-6/models"

type TaskRepository interface {
	CreateTask(task *models.Task) (*models.Task, error)
	GetTask(id int) (*models.Task, error)
	GetTasks() ([]*models.Task, error)
	UpdateTask(task *models.Task) (*models.Task, error)
	DeleteTask(id int) error
}

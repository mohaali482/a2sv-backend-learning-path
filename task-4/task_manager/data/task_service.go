package data

import (
	"github.com/mohaali482/a2sv-backend-learning-path/task-4/models"
)

type TaskUseCase interface {
	GetAllTasks() []*models.Task
	GetTaskById(id int) (*models.Task, error)
	UpdateTask(id int, task models.Task) (*models.Task, error)
	DeleteTask(id int) error
	CreateTask(task models.Task) error
}

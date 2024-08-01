package data

import (
	"context"
	"errors"
	"sync"

	"github.com/mohaali482/a2sv-backend-learning-path/task-5/models"
)

type TaskUseCase interface {
	GetAllTasks(ctx context.Context) []*models.Task
	GetTaskById(ctx context.Context, id int) (*models.Task, error)
	UpdateTask(ctx context.Context, id int, task models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, id int) error
	CreateTask(ctx context.Context, task models.Task) models.Task
}

type TaskService struct {
	nextId int
	tasks  map[int]*models.Task
	mu     sync.RWMutex
}

func NewTaskService() *TaskService {
	return &TaskService{
		nextId: 1,
		tasks:  make(map[int]*models.Task),
		mu:     sync.RWMutex{},
	}
}

func (s *TaskService) GetAllTasks(ctx context.Context) []*models.Task {
	s.mu.RLock()

	tasks := make([]*models.Task, 0)
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	s.mu.RUnlock()

	return tasks
}

func (s *TaskService) GetTaskById(ctx context.Context, id int) (*models.Task, error) {
	s.mu.RLock()
	task, ok := s.tasks[id]
	s.mu.RUnlock()

	if !ok {
		return &models.Task{}, errors.New("task not found")
	}

	return task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id int, task models.Task) (models.Task, error) {
	oldTask, err := s.GetTaskById(ctx, id)

	if err != nil {
		return models.Task{}, err
	}

	s.mu.Lock()

	oldTask.Title = task.Title
	oldTask.Done = task.Done

	s.mu.Unlock()

	return *oldTask, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	_, err := s.GetTaskById(ctx, id)
	if err != nil {
		return err
	}

	s.mu.Lock()

	delete(s.tasks, id)

	s.mu.Unlock()

	return nil
}

func (s *TaskService) CreateTask(ctx context.Context, task models.Task) models.Task {
	newTask := &models.Task{
		Title: task.Title,
		Done:  false,
	}

	s.mu.Lock()

	newTask.Id = s.nextId
	s.nextId++
	s.tasks[newTask.Id] = newTask

	s.mu.Unlock()

	return *newTask
}

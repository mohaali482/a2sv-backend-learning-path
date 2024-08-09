package usecases

import (
	"context"

	domain "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Domain"
	repositories "github.com/mohaali482/a2sv-backend-learning-path/task-7/task-manager/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
	Repository  repositories.TaskRepository
	UserUsecase UserUsecase
}

func (s *TaskUsecase) GetAllTasks(ctx context.Context) []*domain.Task {
	return s.Repository.GetAllTasks(ctx)
}

func (s *TaskUsecase) GetUserTasks(ctx context.Context, userId string) []*domain.Task {
	return s.Repository.GetUserTasks(ctx, userId)
}

func (s *TaskUsecase) GetUserTaskById(ctx context.Context, id string, userId string) (*domain.Task, error) {
	return s.Repository.GetUserTaskById(ctx, id, userId)
}

func (s *TaskUsecase) GetTaskById(ctx context.Context, id string) (*domain.Task, error) {
	return s.Repository.GetTaskById(ctx, id)
}

func (s *TaskUsecase) UpdateTask(ctx context.Context, id string, task domain.Task) (domain.Task, error) {
	err := s.UserUsecase.IsValidId(ctx, task.UserId.Hex())
	if err != nil {
		return domain.Task{}, err
	}

	err = s.Repository.UpdateTask(ctx, id, task)

	if err != nil {
		return domain.Task{}, err
	}

	updatedTask, err := s.Repository.GetTaskById(ctx, id)
	if err != nil {
		return domain.Task{}, err
	}

	return *updatedTask, err
}

func (s *TaskUsecase) DeleteTask(ctx context.Context, id string) error {
	return s.Repository.DeleteTask(ctx, id)
}

func (s *TaskUsecase) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	err := s.UserUsecase.IsValidId(ctx, task.UserId.Hex())
	if err != nil {
		return domain.Task{}, err
	}
	taskId, err := s.Repository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	taskObjectId, _ := primitive.ObjectIDFromHex(taskId)
	task.Id = taskObjectId

	return task, nil
}

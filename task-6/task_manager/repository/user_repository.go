package repository

import "github.com/mohaali482/a2sv-backend-learning-path/task-6/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

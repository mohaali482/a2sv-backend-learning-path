package domain

import "errors"

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyPromoted = errors.New("user is already promoted")
var ErrUniqueUsername = errors.New("user with this username already exists")
var ErrInvalidUserId = errors.New("invalid user id")

var ErrTaskNotFound = errors.New("task not found")
var ErrInvalidTaskId = errors.New("invalid task id")

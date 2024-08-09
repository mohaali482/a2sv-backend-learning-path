package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	admin = iota
	user
)

type User struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Role     int    `json:"role,omitempty" bson:"role,omitempty"`
}

func (u *User) GetRole() string {
	switch u.Role {
	case admin:
		return "admin"
	case user:
		return "user"
	default:
		return "user"
	}
}

type Task struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	DateTime    time.Time          `json:"datetime,omitempty" bson:"datetime,omitempty"`
	Done        bool               `json:"done,omitempty" bson:"done,omitempty"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	admin = iota
	user
)

type User struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
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

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"user_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DateTime    time.Time          `json:"datetime"`
	Done        bool               `json:"done"`
}

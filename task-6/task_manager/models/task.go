package models

import "time"

type Task struct {
	Id     string `json:"id" bson:"_id,omitempty"`
	UserId string `json:"user_id"`

	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateTime    time.Time `json:"datetime"`
	Done        bool      `json:"done"`
}

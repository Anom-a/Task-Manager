package models

import "time"

type Task struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"due_date,omitempty" bson:"due_date,omitempty"`
	Status      string    `json:"status" bson:"status"`
}

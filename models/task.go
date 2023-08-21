package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DueDate     string             `json:"due_date"`
	Priority    string             `json:"priority"`
	Tags        []string           `json:"tags"`
	Subtasks    []Subtask          `json:"subtasks"`
	Attachments []Attachment       `json:"attachments"`
	Done        bool               `json:"done"`
}

type Subtask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type Attachment struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

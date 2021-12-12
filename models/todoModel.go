package models

import (
	"time"
)

// Todo struct
// This struct models a regular Todo
type Todo struct {
	Id string `json:"id"`
	// CreatedAt
	// This field is the time in the UTC
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt
	// This field is the time in the UTC
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
}

type TodoInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed" binding:"required"`
}

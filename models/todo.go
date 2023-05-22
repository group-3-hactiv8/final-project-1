package models

import "gorm.io/gorm"

// https://jsonplaceholder.typicode.com/todos

// Todo represents a model for a todo
type Todo struct {
	gorm.Model
	UserID    uint   `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

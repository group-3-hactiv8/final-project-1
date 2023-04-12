package models

import "gorm.io/gorm"

// https://jsonplaceholder.typicode.com/todos
type Todo struct {
	gorm.Model
	UserID    uint   `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

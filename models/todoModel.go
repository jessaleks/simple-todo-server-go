package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t *Todo) CreateTodo(db *gorm.DB) {
	db.Create(&t)
}

func (t *Todo) GetTodo(db *gorm.DB, id uint) (Todo, error) {
	var todo Todo
	result := db.First(&todo, id)
	if result.Error != nil {
		return todo, result.Error
	}
	return todo, nil
}

func (t *Todo) GetTodos(db *gorm.DB, page int, pageSize int) ([]Todo, error) {
	if pageSize <= 0 {
		pageSize = 10
	}
	var todos []Todo
	result := db.Offset(page * pageSize).Limit(10).Find(&todos)

	if result.Error != nil {
		return todos, result.Error
	}
	return todos, nil
}

func (t *Todo) UpdateTodo(db *gorm.DB) (Todo, error) {
	var todo Todo
	db.Save(&todo)
	return todo, nil
}

func (t *Todo) DeleteTodo(db *gorm.DB) (Todo, error) {
	var todo Todo
	result := db.Delete(&todo)

	if result.Error != nil {
		return todo, result.Error
	}
	return todo, nil
}

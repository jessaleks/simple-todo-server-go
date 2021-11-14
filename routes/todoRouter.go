package routes

import (
	"github.com/jessaleks/simple-todo-server-go/models"
)

func getTodos(c *Context) {
	todos, err := models.Todo.GetTodos()
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, todos)
}
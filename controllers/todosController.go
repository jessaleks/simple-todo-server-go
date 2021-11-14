package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jessaleks/simple-todo-server-go/models"
	"github.com/jessaleks/simple-todo-server-go/utils"
)

//todoRouter.GET("/", GetTodos)
//todoRouter.GET("/:id", GetTodo)
//todoRouter.POST("/", CreateTodo)
//todoRouter.UPDATE("/:id", UpdateTodo)
//todoRouter.DELETE("/:id", DeleteTodo)

// GET /todos
// GetTodos gets an array of Todos with a limit of limit and an offset of offset
func GetTodos(c *gin.Context, limit int, offset int) {
	var todos []models.Todo
	// Get the todos from database
	err := utils.DB.Limit(limit).Offset(offset).Find(&todos).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	// Return the todos in the response
	c.JSON(200, todos)
}

// GET /todos/:id
// GetTodo gets a Todo with the id provided
func GetTodo(c *gin.Context) {
	// Get the id from the url
	id := c.Params.ByName("id")
	// Get the todo from database
	var todo models.Todo
	err := utils.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	// Return the todo in the response
	c.JSON(200, todo)
}

// POST /todos
// CreateTodo creates a new Todo
func CreateTodo(c *gin.Context) {
	var input models.TodoInput
	// Checking if the input is valid
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, err)
		return
	}

	book := models.Todo{Title: input.Title, Description: input.Description}
	// Insert the todo in the database
	err := utils.DB.Create(&book).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// Time to send back the todo in the response
	c.JSON(http.StatusCreated, book)

}
func UpdateTodo(c *gin.Context) {

}
func DeleteTodo(c *gin.Context) {

}

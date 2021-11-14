package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jessaleks/simple-todo-server-go/models"
	"github.com/jessaleks/simple-todo-server-go/utils"
)

//todoRouter.GET("/", GetTodos)
//todoRouter.GET("/:id", GetTodo)
//todoRouter.POST("/", CreateTodo)
//todoRouter.UPDATE("/:id", UpdateTodo)
//todoRouter.DELETE("/:id", DeleteTodo)

// GetTodos GET /
// Gets an array of Todos with a limit of limit and an offset of offset
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	// Get the todos from database
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = utils.DB.Limit(limit).Offset(offset).Find(&todos).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	// Return the todos in the response
	c.JSON(200, todos)
}

// GetTodo GET /:id
// Gets a Todo with the id provided
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

// CreateTodo POST /
// Creates a new Todo
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

// UpdateTodo UPDATE /:id
// Updates a Todo with the id provided
func UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	// Get the todo from database
	err := utils.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	// Update the todo
	err = c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(500, err)
		return
	}
	// Save the todo in the database
	err = utils.DB.Save(&todo).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	// Return the updated todo in the response
	c.JSON(200, todo)
}

// DeleteTodo DELETE /:id
// Deletes a Todo from the DB based on the id provided
func DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	// Get the todo from database
	err := utils.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	// Delete the todo from the database
	err = utils.DB.Delete(&todo).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	// Return the deleted todo in the response
	c.JSON(200, todo)
}

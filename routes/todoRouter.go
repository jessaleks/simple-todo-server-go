package routes

import (
	_ "github.com/gin-gonic/gin"
	"github.com/jessaleks/simple-todo-server-go/controllers"
)

r := gin.New()

todoRouter := r.Group("/todos") {
	todoRouter.GET("/", controllers.GetTodos)
	todoRouter.GET("/:id", controllers.GetTodo)
	todoRouter.POST("/", controllers.CreateTodo)
	todoRouter.UPDATE("/:id", controllers.UpdateTodo)
	todoRouter.DELETE("/:id", controllers.DeleteTodo)
}



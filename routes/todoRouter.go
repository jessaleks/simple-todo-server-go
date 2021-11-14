package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jessaleks/simple-todo-server-go/controllers"
)

func InitializeRoutes(router *gin.Engine) *gin.Engine {
	router.Group("/v1")
	{
		router.GET("todo", controllers.GetTodos)
		router.GET("todo/:id", controllers.GetTodo)
		router.POST("todo", controllers.CreateTodo)
		router.PATCH("todo/:id", controllers.UpdateTodo)
		router.DELETE("todo/:id", controllers.DeleteTodo)

	}
	return router
}

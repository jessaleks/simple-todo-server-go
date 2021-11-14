package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jessaleks/simple-todo-server-go/routes"
	"github.com/jessaleks/simple-todo-server-go/utils"
)

func main() {
	r := gin.Default()

	utils.ConnectToTheDatabase()
	routes.InitializeRoutes(r)
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]string{"Hello": "Hello"})
	})
	err := r.Run()
	if err != nil {
		log.Panicln(err)
	}

}

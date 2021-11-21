package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jessaleks/simple-todo-server-go/routes"
	"github.com/jessaleks/simple-todo-server-go/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Connecting to the DB
	_, err := utils.ConnectToTheDatabase()
	if err != nil {
		panic(err)
	}
	routes.InitializeRoutes(r)

	// The index route, just to see if the server really works lol
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]string{"Hello": "Hello"})
	})

	// The swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	utils.SetUpMonitoring()
	r.GET("/metrics", utils.PrometheusHandler())

	// running the app
	err = r.Run()
	if err != nil {
		log.Panicln(err)
	}

}

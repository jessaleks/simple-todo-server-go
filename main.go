package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jessaleks/simple-todo-server-go/utils"
	"log"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	utils.ConnectToTheDatabase()

	err := r.Run()
	if err != nil {
		log.Println(err)
	}
	println("Hello, world!")
}

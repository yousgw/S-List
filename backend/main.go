package main

import (
	todo "./controller/todo.go"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	todoRouter := router.Group("/todo")
	{
		todoRouter.GET("/get", todo.outAll)
		todoRouter.POST("/add", controller.in)
	}

	router.Run(":3000")

}

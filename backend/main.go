package main

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	todoRouter := router.Group("/todo")
	{
		todoRouter.GET("/get", controller.outAll)
		todoRouter.POST("/add", controller.in)
	}

	router.Run(":3000")

}

package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	todoRouter := router.Group("/todo")
	{
		v1 := todoRouter.Group("/v1")
		{
			v1.GET("/get", controller.OutAll)
			v1.POST("/add", controller.In)
		}

	}

	router.Run(":3000")

}

package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	todoRouter := router.Group("/todo")
	{
			todoRouter.GET("/get", controller.OutAll)
			todoRouter.POST("/add", controller.In)
			todoRouter.GET("/", controller.Test)

	}

	router.Run(":3000")

}

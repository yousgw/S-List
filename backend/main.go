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
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))

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

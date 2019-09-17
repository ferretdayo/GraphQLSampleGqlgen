package infrastructure

import (
	"github.com/GraphQLSample/src/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userController := controllers.NewUserController()
	router.GET("/users", func(c *gin.Context) { userController.Create(c) })
	router.Run() // listen and serve on 0.0.0.0:8080
}

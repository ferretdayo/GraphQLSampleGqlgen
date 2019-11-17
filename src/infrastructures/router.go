package infrastructures

import (
	handler "github.com/GraphQLSample/src/infrastructures/handlers"
	"github.com/GraphQLSample/src/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	graphQl := handler.NewGraphQL()
	router.GET("/graphql", graphQl.Handler())
	router.POST("/graphql", graphQl.Handler())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userController := controllers.NewUserController()
	// router.GET("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Create(c) })
	router.Run() // listen and serve on 0.0.0.0:8080
}

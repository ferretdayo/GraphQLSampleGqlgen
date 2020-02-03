package infrastructures

import (
	"github.com/GraphQLSample/src/infrastructures/db"
	handler "github.com/GraphQLSample/src/infrastructures/handlers"
	"github.com/GraphQLSample/src/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	mysql := db.NewMysql()
	db := mysql.Open()
	graphQl := handler.NewGraphQL()
	router.GET("/graphql", graphQl.Handler())
	router.POST("/graphql", graphQl.Handler())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userController := controllers.NewUserController(db)
	router.GET("/users", func(c *gin.Context) { userController.Create(c) })
	router.Run() // listen and serve on 0.0.0.0:8080
}

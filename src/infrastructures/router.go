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
	router.GET("/users", func(c *gin.Context) { userController.GetUsers(c) })
	router.GET("/users/:user_id", func(c *gin.Context) { userController.GetUser(c) })
	// router.POST("/users", func(c *gin.Context) { userController.Create(c) })

	hobbyController := controllers.NewHobbyController(db)
	router.GET("/masters/hobbies", func(c *gin.Context) { hobbyController.GetHobbies(c) })

	router.Run() // listen and serve on 0.0.0.0:8080
}

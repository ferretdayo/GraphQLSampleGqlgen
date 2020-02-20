package infrastructures

import (
	"github.com/GraphQLSampleGqlgen/src/infrastructures/db"
	"github.com/GraphQLSampleGqlgen/src/infrastructures/handlers"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	mysql := db.NewMysql()
	db := mysql.Open()
	graphQl := handler.NewGraphQL(db)
	router.GET("/graphql", graphQl.Handler())
	router.POST("/graphql", graphQl.Handler())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}

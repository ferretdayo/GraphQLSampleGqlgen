package handler

import (
	"github.com/GraphQLSample/src/usecases"
	"github.com/GraphQLSample/src/usecases/queries"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type GraphQL struct {
	UserUsecase  *usecases.UserUsecase
	HobbyUsecase *usecases.HobbyUsecase
	UserQuery    *queries.UserQuery
	HobbyQuery   *queries.HobbyQuery
}

func NewGraphQL() *GraphQL {
	return &GraphQL{
		UserUsecase:  &usecases.UserUsecase{},
		HobbyUsecase: &usecases.HobbyUsecase{},
		UserQuery:    &queries.UserQuery{},
		HobbyQuery:   &queries.HobbyQuery{},
	}
}

func (graphQl *GraphQL) Handler() gin.HandlerFunc {
	// fields := graphql.Fields{
	// 	"UserQuery": &graphql.Field{
	// 		Name:   "User",graphQl.HobbyQuery.CreateHobbyQuery()
	// 		Fields: graphQl.UserQuery.CreateUserQuery(),
	// 	},
	// 	"HobbyQuery": &graphql.Field{
	// 		Name:   "Hobby",
	// 		Fields: ,
	// 	},
	// }
	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Root",
			Fields: graphQl.UserQuery.CreateUserQuery(),
		},
		// graphql.ObjectConfig{
		// 	Name:   "Hobby",
		// 	Fields: graphQl.HobbyQuery.CreateHobbyQuery(),
		// },
	)
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

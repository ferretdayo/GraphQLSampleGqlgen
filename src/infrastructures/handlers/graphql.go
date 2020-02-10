package handler

import (
	"github.com/GraphQLSample/src/infrastructures/db"
	"github.com/GraphQLSample/src/interfaces/repositories"
	"github.com/GraphQLSample/src/usecases/queries"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type GraphQL struct {
	UserQuery *queries.UserQuery
	// HobbyQuery   *queries.HobbyQuery
}

func NewGraphQL(db *db.Database) *GraphQL {
	return &GraphQL{
		UserQuery: &queries.UserQuery{
			UserRepository: &repositories.UserRepository{},
			DB:             db,
		},
		// HobbyQuery:   &queries.HobbyQuery{},
	}
}

func (graphQl *GraphQL) Handler() gin.HandlerFunc {
	fields := graphql.Fields{
		"User": graphQl.UserQuery.CreateUserQuery(),
		// "Hobby": graphQl.HobbyQuery.CreateHobbyQuery(),
	}
	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Root",
			Fields: fields,
		},
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

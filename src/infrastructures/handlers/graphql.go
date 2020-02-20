package handler

import (
	"github.com/GraphQLSample/src/infrastructures/db"
	"github.com/GraphQLSample/src/infrastructures/schema"
	"github.com/GraphQLSample/src/interfaces/repositories"
	"github.com/GraphQLSample/src/usecases/resolvers/masters"
	"github.com/GraphQLSample/src/usecases/resolvers/users"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type GraphQL struct {
	DB *db.Database
}

func NewGraphQL(db *db.Database) *GraphQL {
	return &GraphQL{
		DB: db,
	}
}

func (graphQl *GraphQL) Handler() gin.HandlerFunc {
	rootSchema := schema.NewRootSchema(
		users.UserResolver{
			UserRepository:       &repositories.UserRepository{},
			UserDetailRepository: &repositories.UserDetailRepository{},
			DB:                   graphQl.DB,
		},
		masters.HobbyResolver{
			HobbyRepository: &repositories.HobbyRepository{},
			DB:              graphQl.DB,
		})
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootSchema.Query(),
		Mutation: rootSchema.Mutation(),
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

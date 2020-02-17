package schema

import (
	resolvers "github.com/GraphQLSample/src/interfaces/resolvers/users"
	"github.com/GraphQLSample/src/types"
	"github.com/graphql-go/graphql"
)

type RootSchema struct {
	UserResolver resolvers.UserResolver
}

func NewRootSchema(resolver resolvers.UserResolver) *RootSchema {
	return &RootSchema{
		UserResolver: resolver,
	}
}

func (schema *RootSchema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"User": &graphql.Field{
				Type: types.UserObjectType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: schema.UserResolver.GetUserByID,
			},
		},
	}
	return graphql.NewObject(objectConfig)
}

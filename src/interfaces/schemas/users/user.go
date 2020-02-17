package schemas

import (
	resolvers "github.com/GraphQLSample/src/interfaces/resolvers/users"
	"github.com/GraphQLSample/src/types"
	"github.com/graphql-go/graphql"
)

type UserSchema struct {
	UserResolver resolvers.UserResolver
}

func NewUserSchema(resolver resolvers.UserResolver) *UserSchema {
	return &UserSchema{
		UserResolver: resolver,
	}
}

func (schema *UserSchema) Query() *graphql.Object {
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

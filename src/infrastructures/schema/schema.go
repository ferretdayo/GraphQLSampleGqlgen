package schema

import (
	"github.com/GraphQLSample/src/types"
	"github.com/GraphQLSample/src/usecases/resolvers/masters"
	"github.com/GraphQLSample/src/usecases/resolvers/users"
	"github.com/graphql-go/graphql"
)

type RootSchema struct {
	UserResolver  users.UserResolver
	HobbyResolver masters.HobbyResolver
}

func NewRootSchema(userResolver users.UserResolver, hobbyResolver masters.HobbyResolver) *RootSchema {
	return &RootSchema{
		UserResolver:  userResolver,
		HobbyResolver: hobbyResolver,
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
			"UserList": &graphql.Field{
				Type:    graphql.NewList(types.UserObjectType),
				Resolve: schema.UserResolver.GetList,
			},
			"HobbyList": &graphql.Field{
				Type:    graphql.NewList(types.MasterObjectType),
				Resolve: schema.HobbyResolver.GetList,
			},
		},
	}
	return graphql.NewObject(objectConfig)
}

func (schema *RootSchema) Mutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"CreateUser": &graphql.Field{
				Type:    types.UserObjectType,
				Resolve: schema.UserResolver.CreateUser,
			},
		},
	}
	return graphql.NewObject(objectConfig)
}

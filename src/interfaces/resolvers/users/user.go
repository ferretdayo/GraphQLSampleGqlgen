package resolvers

import "github.com/graphql-go/graphql"

type UserResolver interface {
	GetUserByID(params graphql.ResolveParams) (interface{}, error)
}

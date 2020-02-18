package types

import "github.com/graphql-go/graphql"

var MasterObjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Master",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"is_deleted": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
})

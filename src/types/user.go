package types

import (
	"github.com/graphql-go/graphql"
)

var UserObjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"display_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"is_unsubscribed": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"created_at": &graphql.Field{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
		"updated_at": &graphql.Field{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
		"hobby": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "Hobby",
				Fields: graphql.Fields{
					"id": &graphql.Field{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"name": &graphql.Field{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
			}),
		},
	},
})

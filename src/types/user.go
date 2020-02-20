package types

import (
	"errors"
	"github.com/GraphQLSampleGqlgen/src/usecases/ports"
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
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							c, ok := p.Source.(*ports.MasterOutputPort)
							if ok && c.ID != 0 {
								return c.ID, nil
							}
							return nil, errors.New("id error")
						},
					},
					"name": &graphql.Field{
						Type: graphql.NewNonNull(graphql.String),
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							c, ok := p.Source.(*ports.MasterOutputPort)
							if ok && c.Name != "" {
								return c.Name, nil
							}
							return nil, errors.New("name error")
						},
					},
				},
			}),
		},
	},
})

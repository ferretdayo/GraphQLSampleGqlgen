package queries

import (
	"fmt"

	"github.com/GraphQLSample/src/usecases"

	"github.com/GraphQLSample/src/interfaces/entities"
	"github.com/graphql-go/graphql"
)

type HobbyQuery struct {
	Usecase *usecases.HobbyUsecase
}

func (hobbyQuery *HobbyQuery) CreateHobbyQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Hobby",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
			},
		}),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"]
			v, _ := id.(int)
			fmt.Printf("fetching post with id: %d", v)
			u, err := hobbyQuery.Usecase.GetHobby(&entities.Hobby{
				ID:   0,
				Name: "野球",
			})
			if err != nil {
				return nil, err
			}
			return u, nil
		},
	}
}

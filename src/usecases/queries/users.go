package queries

import (
	"fmt"

	"github.com/GraphQLSample/src/usecases"

	"github.com/GraphQLSample/src/interfaces/entities"
	"github.com/graphql-go/graphql"
)

type UserQuery struct {
	HobbyUsecase *usecases.HobbyUsecase
	Usecase      *usecases.UserUsecase
	HobbyQuery   *HobbyQuery
}

func (userQuery *UserQuery) CreateUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"NickName": &graphql.Field{
					Type: graphql.String,
				},
				"Old": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),
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
			u, err := userQuery.Usecase.GetUser(&entities.User{
				ID:       0,
				NickName: "AAAAA",
				Old:      29,
			})
			if err != nil {
				return nil, err
			}
			return u, nil
		},
	}
}

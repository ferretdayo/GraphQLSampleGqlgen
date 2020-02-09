package queries

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

type UserQuery struct {
}

func (userQuery *UserQuery) CreateUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"IsUnsubscribed": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
				"CreatedAt": &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"UpdatedAt": &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
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
			// u, err := userQuery.Usecase.GetUser(&entities.User{
			// 	ID:       0,
			// 	NickName: "AAAAA",
			// 	Old:      29,
			// })
			// if err != nil {
			// 	return nil, err
			// }
			// return u, nil
			return nil, nil
		},
	}
}

// func (userQuery *UserQuery) CreateHobbyQuery() *graphql.Field {
// 	return &graphql.Field{
// 		Type: graphql.NewObject(graphql.ObjectConfig{
// 			Name: "Hobby",
// 			Fields: graphql.Fields{
// 				"ID": &graphql.Field{
// 					Type: graphql.NewNonNull(graphql.Int),
// 				},
// 				"Name": &graphql.Field{
// 					Type: graphql.String,
// 				},
// 			},
// 		}),
// 		Args: graphql.FieldConfigArgument{
// 			"id": &graphql.ArgumentConfig{
// 				Type: graphql.NewNonNull(graphql.Int),
// 			},
// 		},
// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 			id := p.Args["id"]
// 			v, _ := id.(int)
// 			fmt.Printf("fetching post with id: %d", v)
// 			u, err := hobbyQuery.Usecase.GetHobby(&entities.Hobby{
// 				ID:   0,
// 				Name: "野球",
// 			})
// 			if err != nil {
// 				return nil, err
// 			}
// 			return u, nil
// 		},
// 	}
// }

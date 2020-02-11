package queries

import (
	"github.com/GraphQLSample/src/infrastructures/db"
	"github.com/GraphQLSample/src/usecases/repositories"

	"github.com/graphql-go/graphql"
)

type UserQuery struct {
	UserRepository repositories.UserRepository
	DB             *db.Database
}

func (query *UserQuery) CreateUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"DisplayID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
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
			userID := p.Args["id"].(int)
			user, err := query.UserRepository.SelectByUserID(query.DB.MainDB.ReadReplica, uint(userID))
			if err != nil {
				return nil, err
			}
			return user, nil
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

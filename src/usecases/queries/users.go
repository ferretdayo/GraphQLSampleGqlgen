package queries

import (
	"fmt"

	"github.com/GraphQLSample/src/entities"
	"github.com/GraphQLSample/src/usecases/ports"
	"github.com/GraphQLSample/src/usecases/users"
	"github.com/graphql-go/graphql"
)

type UserQuery struct {
	Usecase *users.UserUsecase
}

var userType = graphql.NewObject(graphql.ObjectConfig{
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
		"hobby": createHobbyQuery(),
	},
})

var hobbyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Hobby",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

func (query *UserQuery) CreateUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userID := p.Args["id"].(int)
			input := &ports.UserInputPort{
				UserID: uint(userID),
			}
			user, err := query.Usecase.GetUser(input)
			if err != nil {
				return nil, err
			}
			return user, nil
		},
	}
}

func (query *UserQuery) CreateUserListQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(userType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			users, err := query.Usecase.GetUsers()
			if err != nil {
				return nil, err
			}
			return users, nil
		},
	}
}

func createHobbyQuery() *graphql.Field {
	return &graphql.Field{
		Type: hobbyType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Printf("p.source: %v\n", p.Source.(*ports.UserOutputPort))
			// userID := p.Args["id"].(int)
			// input := &ports.UserInputPort{
			// 	UserID: uint(userID),
			// }
			// user, err := query.Usecase.GetUser(input)
			// if err != nil {
			// 	return nil, err
			// }
			return entities.Hobby{
				ID:   1,
				Name: "野球",
			}, nil
		},
	}
}

func (query *UserQuery) createUserDetailQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Detail",
			Fields: graphql.Fields{
				"UserID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"Nickname": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"Birthday": &graphql.Field{
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
			id := p.Args["id"].(int)
			userDetail, err := query.Usecase.UserDetailRepository.SelectByUserID(query.Usecase.DB.MainDB.ReadReplica, uint(id))
			if err != nil {
				return nil, err
			}
			return userDetail, nil
		},
	}
}

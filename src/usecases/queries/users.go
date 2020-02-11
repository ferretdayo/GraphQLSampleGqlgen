package queries

import (
	"github.com/GraphQLSample/src/usecases/ports"
	"github.com/GraphQLSample/src/usecases/users"
	"github.com/graphql-go/graphql"
)

type UserQuery struct {
	Usecase *users.UserUsecase
}

func (query *UserQuery) CreateUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.ID),
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
				"Detail": query.CreateUserDetailQuery(),
			},
		}),
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

func (query *UserQuery) CreateUserDetailQuery() *graphql.Field {
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

package resolvers

import (
	"errors"

	"github.com/GraphQLSample/src/infrastructures/db"
	"github.com/GraphQLSample/src/usecases/ports"
	"github.com/GraphQLSample/src/usecases/repositories"
	"github.com/graphql-go/graphql"
)

type UserResolver struct {
	UserRepository       repositories.UserRepository
	UserDetailRepository repositories.UserDetailRepository
	DB                   *db.Database
}

func (resolver *UserResolver) GetUserByID(params graphql.ResolveParams) (interface{}, error) {
	userID := params.Args["id"].(int)
	input := &ports.UserInputPort{
		UserID: uint(userID),
	}

	user, err := resolver.UserRepository.SelectByUserID(resolver.DB.MainDB.ReadReplica, input.UserID)
	if err != nil {
		return nil, errors.New("something wrong.")
	}

	outputPort := &ports.UserOutputPort{
		ID:             user.ID,
		DisplayID:      user.DisplayID,
		IsUnsubscribed: user.IsUnsubscribed,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
	return outputPort, nil
}

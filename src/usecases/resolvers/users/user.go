package users

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/GraphQLSample/src/entities"
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
		Hobby: &ports.MasterOutputPort{
			ID:   user.UserDetail.Hobby.ID,
			Name: user.UserDetail.Hobby.Name,
		},
	}
	return outputPort, nil
}

func (resolver *UserResolver) GetList(params graphql.ResolveParams) (interface{}, error) {
	users, err := resolver.UserRepository.Select(resolver.DB.MainDB.ReadReplica)
	if err != nil {
		return nil, errors.New("something wrong.")
	}

	var output []ports.UserOutputPort
	for _, user := range users {
		output = append(output, ports.UserOutputPort{
			ID:             user.ID,
			DisplayID:      user.DisplayID,
			IsUnsubscribed: user.IsUnsubscribed,
			CreatedAt:      user.CreatedAt,
			UpdatedAt:      user.UpdatedAt,
		})
	}
	return output, nil
}

func (resolver *UserResolver) CreateUser(params graphql.ResolveParams) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	displayID := fmt.Sprintf("%010d", rand.Int31n(math.MaxInt32))

	user := &entities.User{
		DisplayID: displayID,
	}
	if err := resolver.UserRepository.Insert(resolver.DB.MainDB.Master, user); err != nil {
		return nil, err
	}

	output := &ports.UserOutputPort{
		ID:             user.ID,
		DisplayID:      user.DisplayID,
		IsUnsubscribed: user.IsUnsubscribed,
	}
	fmt.Printf("%v", output)
	return output, nil
}

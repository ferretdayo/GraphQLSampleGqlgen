package users

import (
	"errors"

	"github.com/GraphQLSample/src/infrastructures/db"
	"github.com/GraphQLSample/src/usecases/ports"
	"github.com/GraphQLSample/src/usecases/repositories"
)

type UserUsecase struct {
	UserRepository repositories.UserRepository
	DB             *db.Database
}

func (usecase *UserUsecase) GetUsers() (*ports.UsersOutputPort, error) {
	users, err := usecase.UserRepository.Select(usecase.DB.MainDB.ReadReplica)
	if err != nil {
		return nil, errors.New("something wrong.")
	}
	output := &ports.UsersOutputPort{
		Users: users,
	}
	return output, nil
}

func (usecase *UserUsecase) GetUser(input *ports.UserInputPort) (*ports.UserOutputPort, error) {
	user, err := usecase.UserRepository.SelectByUserID(usecase.DB.MainDB.ReadReplica, input.UserID)
	if err != nil {
		return nil, errors.New("something wrong.")
	}

	outputPort := &ports.UserOutputPort{
		ID:             user.ID,
		IsUnsubscribed: user.IsUnsubscribed,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
	return outputPort, nil
}

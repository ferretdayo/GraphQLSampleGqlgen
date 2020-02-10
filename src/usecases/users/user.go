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
)

type UserUsecase struct {
	UserRepository repositories.UserRepository
	DB             *db.Database
}

func (usecase *UserUsecase) CreateUser() (*ports.UserOutputPort, error) {
	rand.Seed(time.Now().UnixNano())
	displayID := fmt.Sprintf("%010d\n", rand.Int31n(math.MaxInt32))

	user := &entities.User{
		DisplayID: displayID,
	}
	if err := usecase.UserRepository.Insert(usecase.DB.MainDB.Master, user); err != nil {
		return nil, err
	}

	output := &ports.UserOutputPort{
		ID: user.ID,
	}
	return output, nil
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

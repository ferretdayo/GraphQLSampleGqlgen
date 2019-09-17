package usecases

import (
	"errors"

	"github.com/GraphQLSample/src/interfaces/entities"
)

type UserUsecase struct{}

func (usecase *UserUsecase) GetUser(u *entities.User) (*entities.User, error) {
	if u == nil {
		return nil, errors.New("something wrong.")
	}
	return u, nil
}

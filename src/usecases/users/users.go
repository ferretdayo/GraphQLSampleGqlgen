package users

import (
	"errors"

	"github.com/GraphQLSample/src/usecases/repositories"

	"github.com/GraphQLSample/src/entities"
)

type UserUsecase struct {
	UserRepository repositories.UserRepository
	DB             *db.DB
}

func (usecase *UserUsecase) GetUser(u *entities.User) (*entities.User, error) {
	user := usecase.UserRepository.Select()
	if u == nil {
		return nil, errors.New("something wrong.")
	}
	return u, nil
}

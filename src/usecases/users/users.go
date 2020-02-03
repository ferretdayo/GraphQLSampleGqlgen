package users

import (
	"errors"

	"github.com/GraphQLSample/src/infrastructures/db"
	"github.com/GraphQLSample/src/usecases/repositories"

	"github.com/GraphQLSample/src/entities"
)

type UserUsecase struct {
	UserRepository repositories.UserRepository
	DB             *db.Database
}

func (usecase *UserUsecase) GetUser(u *entities.User) ([]entities.User, error) {
	users, err := usecase.UserRepository.Select(usecase.DB.MainDB.ReadReplica)
	if err != nil {
		return nil, errors.New("something wrong.")
	}
	return users, nil
}

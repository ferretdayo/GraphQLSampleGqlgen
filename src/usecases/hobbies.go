package usecases

import (
	"errors"

	"github.com/GraphQLSample/src/interfaces/entities"
)

type HobbyUsecase struct{}

func (usecase *HobbyUsecase) GetHobby(u *entities.Hobby) (*entities.Hobby, error) {
	if u == nil {
		return nil, errors.New("something wrong.")
	}
	return u, nil
}

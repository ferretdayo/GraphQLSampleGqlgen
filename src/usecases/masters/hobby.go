package masters

import (
	"errors"

	"github.com/GraphQLSample/src/infrastructures/db"
	"github.com/GraphQLSample/src/usecases/ports"
	"github.com/GraphQLSample/src/usecases/repositories"
)

type HobbyUsecase struct {
	HobbyRepository repositories.HobbyRepository
	DB              *db.Database
}

func (usecase *HobbyUsecase) GetHobbies() ([]ports.MasterOutputPort, error) {
	hobbies, err := usecase.HobbyRepository.Select(usecase.DB.MainDB.ReadReplica)
	if err != nil {
		return nil, errors.New("something wrong.")
	}
	var output []ports.MasterOutputPort
	for _, hobby := range hobbies {
		output = append(output, ports.MasterOutputPort{
			ID:       hobby.ID,
			Name:     hobby.Name,
			IsDelete: hobby.IsDelete,
		})
	}
	return output, nil
}

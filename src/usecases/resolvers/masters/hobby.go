package masters

import (
	"errors"

	"github.com/GraphQLSample/src/infrastructures/db"
	"github.com/GraphQLSample/src/usecases/ports"
	"github.com/GraphQLSample/src/usecases/repositories"
	"github.com/graphql-go/graphql"
)

type HobbyResolver struct {
	HobbyRepository       repositories.HobbyRepository
	DB                   *db.Database
}

func (resolver *HobbyResolver) GetList(params graphql.ResolveParams) (interface{}, error) {
	hobbies, err := resolver.HobbyRepository.Select(resolver.DB.MainDB.ReadReplica)
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

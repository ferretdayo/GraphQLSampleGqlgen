package repositories

import (
	"github.com/GraphQLSampleGqlgen/src/entities"
	"github.com/jinzhu/gorm"
)

type HobbyRepository interface {
	Select(*gorm.DB) ([]entities.Hobby, error)
}

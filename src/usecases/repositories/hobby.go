package repositories

import (
	"github.com/GraphQLSample/src/entities"
	"github.com/jinzhu/gorm"
)

type HobbyRepository interface {
	Select(*gorm.DB) ([]entities.Hobby, error)
}

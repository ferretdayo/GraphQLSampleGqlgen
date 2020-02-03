package repositories

import (
	"github.com/GraphQLSample/src/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Select(*gorm.DB) ([]entities.User, error)
}

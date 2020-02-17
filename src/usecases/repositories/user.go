package repositories

import (
	"github.com/GraphQLSample/src/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(*gorm.DB, *entities.User) error
	Select(*gorm.DB) ([]entities.User, error)
	SelectByUserID(*gorm.DB, uint) (*entities.User, error)
}

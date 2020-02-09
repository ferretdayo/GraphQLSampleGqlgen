package repositories

import (
	"github.com/GraphQLSample/src/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Select(*gorm.DB) ([]entities.User, error)
	SelectByUserID(*gorm.DB, uint) (*entities.User, error)
}

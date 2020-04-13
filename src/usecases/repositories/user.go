package repositories

import (
	"github.com/GraphQLSampleGqlgen/src/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(*gorm.DB, *entities.User) error
	SelectByTokenID(*gorm.DB, string) (*entities.User, error)
	Select(*gorm.DB) ([]entities.User, error)
	SelectByUserID(*gorm.DB, uint) (*entities.User, error)
}

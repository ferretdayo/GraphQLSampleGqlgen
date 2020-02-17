package repositories

import (
	"github.com/GraphQLSample/src/entities"
	"github.com/jinzhu/gorm"
)

type UserDetailRepository interface {
	SelectByUserID(*gorm.DB, uint) (*entities.UserDetail, error)
}

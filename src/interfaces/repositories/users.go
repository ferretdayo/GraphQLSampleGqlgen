package repositories

import (
	"github.com/GraphQLSample/src/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
}

func (repository *UserRepository) Select(db *gorm.DB) ([]entities.User, error) {
	var users []entities.User
	err := db.Find(&users).Error
	return users, err
}

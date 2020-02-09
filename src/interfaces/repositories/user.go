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

func (repository *UserRepository) SelectByUserID(db *gorm.DB, userID uint) (*entities.User, error) {
	var user entities.User
	err := db.Where("id = ?", userID).First(&user).Error
	return &user, err
}

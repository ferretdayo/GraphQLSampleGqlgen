package repositories

import (
	"github.com/GraphQLSample/src/entities"
	"github.com/jinzhu/gorm"
)

type UserDetailRepository struct {
}

func (repository *UserDetailRepository) SelectByUserID(db *gorm.DB, userID uint) (*entities.UserDetail, error) {
	var userDetail entities.UserDetail
	err := db.Where("id = ?", userID).First(&user).Error
	return &userDetail, err
}

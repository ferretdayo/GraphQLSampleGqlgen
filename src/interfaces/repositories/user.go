package repositories

import (
	"github.com/GraphQLSampleGqlgen/src/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
}

func (repository *UserRepository) Insert(db *gorm.DB, user *entities.User) error {
	return db.Save(user).Error
}

func (repository *UserRepository) SelectByTokenID(db *gorm.DB, token string) (*entities.User, error) {
	var user entities.User
	err := db.Model(&user).
		Where("token = ?", token).
		First(&user).Error
	return &user, err
}

func (repository *UserRepository) Select(db *gorm.DB) ([]entities.User, error) {
	var users []entities.User
	err := db.Model(&users).
		Preload("UserTodos").
		Preload("UserDetail").
		Preload("UserDetail.Hobby").
		Find(&users).Error
	return users, err
}

func (repository *UserRepository) SelectByUserID(db *gorm.DB, userID uint) (*entities.User, error) {
	var user entities.User
	err := db.Model(&user).
		Where("id = ?", userID).
		Preload("UserTodos").
		Preload("UserDetail").
		Preload("UserDetail.Hobby").
		First(&user).Error
	return &user, err
}

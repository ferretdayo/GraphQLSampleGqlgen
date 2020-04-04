package repositories

import (
	"github.com/GraphQLSampleGqlgen/src/entities"
	"github.com/jinzhu/gorm"
)

type HobbyRepository struct {
}

func (repository *HobbyRepository) Select(db *gorm.DB) ([]entities.Hobby, error) {
	var hobbies []entities.Hobby
	err := db.Find(&hobbies).Error
	return hobbies, err
}

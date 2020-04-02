package repositories

import (
	"github.com/GraphQLSampleGqlgen/src/entities"
	"github.com/jinzhu/gorm"
)

type UserTodoRepository struct {
}

func (repository *UserTodoRepository) Insert(db *gorm.DB, userTodo *entities.UserTodo) error {
	return db.Save(userTodo).Error
}

func (repository *UserTodoRepository) Select(db *gorm.DB) ([]entities.UserTodo, error) {
	var userTodos []entities.UserTodo
	err := db.Find(&userTodos).
		Preload("User").
		Preload("User.UserDetail").Error
	return userTodos, err
}

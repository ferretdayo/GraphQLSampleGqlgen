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

func (repository *UserTodoRepository) SelectByUserID(db *gorm.DB, userID uint) ([]entities.UserTodo, error) {
	var userTodos []entities.UserTodo
	err := db.Find(&userTodos, "user_id = ?", userID).Error
	return userTodos, err
}

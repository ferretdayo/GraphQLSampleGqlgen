package repositories

import (
	"github.com/GraphQLSampleGqlgen/src/entities"
	"github.com/jinzhu/gorm"
)

type UserTodoRepository interface {
	Insert(*gorm.DB, *entities.UserTodo) error
	SelectByUserID(*gorm.DB, uint) ([]entities.UserTodo, error)
}

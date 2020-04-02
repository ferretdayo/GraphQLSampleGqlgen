package repositories

import (
	"github.com/GraphQLSampleGqlgen/src/entities"
	"github.com/jinzhu/gorm"
)

type UserTodoRepository interface {
	Insert(*gorm.DB, *entities.UserTodo) error
	Select(*gorm.DB) ([]entities.UserTodo, error)
}

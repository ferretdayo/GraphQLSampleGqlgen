// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package resolver

import (
	"github.com/GraphQLSampleGqlgen/src/db"
	"github.com/GraphQLSampleGqlgen/src/usecases/repositories"
)

type Resolver struct {
	UserRepository       repositories.UserRepository
	UserDetailRepository repositories.UserDetailRepository
	UserTodoRepository   repositories.UserTodoRepository
	HobbyRepository      repositories.HobbyRepository
	DB                   *db.Database
}

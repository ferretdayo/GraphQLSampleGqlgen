// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package resolver

import (
	"github.com/GraphQLSampleGqlgen/src/usecases/users"
)

type Resolver struct {
	UserUsecase     *users.UserUsecase
	UserTodoUsecase *users.UserTodoUsecase
}

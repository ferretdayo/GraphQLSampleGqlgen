package middlewares

import (
	"github.com/GraphQLSampleGqlgen/src/db"
	"github.com/GraphQLSampleGqlgen/src/interfaces/api"

	"github.com/GraphQLSampleGqlgen/src/interfaces/repositories"
	"github.com/GraphQLSampleGqlgen/src/usecases/users"
	"github.com/gin-gonic/gin"
)

//UserMiddlewares
type UserMiddleware struct {
	Usecase *users.UserUsecase
}

//NewUserMiddleware
func NewUserMiddleware(db *db.Database) *UserMiddleware {
	return &UserMiddleware{
		Usecase: &users.UserUsecase{
			UserRepository: &repositories.UserRepository{},
			DB:             db,
		},
	}
}

// SetUserToContext ユーザー情報をgin.Contextに格納する
func (handler *UserMiddleware) SetUserToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := handler.Usecase.UserRepository.SelectByTokenID(handler.Usecase.DB.MainDB.ReadReplica, api.GetTokenID(c))
		if err != nil {
			c.JSON(500, api.GetErrorResponse(err))
			c.Abort() //handlerでエラーが起こった場合はそこで終了させること
		}
		api.SetUserID(c, user.ID)
		c.Next()
	}
}

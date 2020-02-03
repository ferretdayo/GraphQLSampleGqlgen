package controllers

import (
	"net/http"

	"github.com/GraphQLSample/src/infrastructures/db"

	"github.com/GraphQLSample/src/entities"
	"github.com/GraphQLSample/src/interfaces/repositories"
	"github.com/GraphQLSample/src/usecases/ports"
	"github.com/GraphQLSample/src/usecases/users"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Usecase *users.UserUsecase
}

func NewUserController(db *db.Database) *UserController {
	return &UserController{
		Usecase: &users.UserUsecase{
			UserRepository: &repositories.UserRepository{},
			DB:             db,
		},
	}
}

func (controller *UserController) Create(c *gin.Context) {
	user := entities.User{
		ID:       1,
		NickName: "aaa",
	}
	c.Bind(&user)
	u, err := controller.Usecase.GetUser(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	outputPort := &ports.UserOutputPort{
		ID:       u[0].ID,
		NickName: u[0].NickName,
		Old:      u[0].Old,
	}
	c.JSON(http.StatusOK, outputPort)
}

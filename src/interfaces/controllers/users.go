package controllers

import (
	"net/http"

	"github.com/GraphQLSample/src/entities"
	"github.com/GraphQLSample/src/usecases/ports"
	"github.com/GraphQLSample/src/usecases/users"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Usecase *users.UserUsecase
}

func NewUserController() *UserController {
	return &UserController{
		Usecase: &users.UserUsecase{},
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
		ID:       u.ID,
		NickName: u.NickName,
		Old:      u.Old,
	}
	c.JSON(http.StatusOK, outputPort)
}

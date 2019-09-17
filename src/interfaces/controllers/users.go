package controllers

import (
	"net/http"

	"github.com/GraphQLSample/src/interfaces/entities"
	"github.com/GraphQLSample/src/usecases"
	"github.com/GraphQLSample/src/usecases/ports"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase usecases.UserUsecase
}

func NewUserController() *UserController {
	return &UserController{}
}

func (controller *UserController) Create(c *gin.Context) {
	user := entities.User{}
	c.Bind(&user)
	u, err := controller.UserUsecase.GetUser(&user)
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

package controllers

import (
	"net/http"
	"strconv"

	"github.com/GraphQLSample/src/usecases/ports"

	"github.com/GraphQLSample/src/infrastructures/db"

	"github.com/GraphQLSample/src/interfaces/repositories"
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
	output, err := controller.Usecase.CreateUser()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, output)
}

func (controller *UserController) GetUsers(c *gin.Context) {
	output, err := controller.Usecase.GetUsers()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, output)
}

func (controller *UserController) GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	input := &ports.UserInputPort{
		UserID: uint(userID),
	}
	output, err := controller.Usecase.GetUser(input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, output)
}

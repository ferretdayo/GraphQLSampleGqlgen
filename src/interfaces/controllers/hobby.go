package controllers

import (
	"net/http"

	"github.com/GraphQLSample/src/infrastructures/db"
	"github.com/GraphQLSample/src/usecases/masters"

	"github.com/GraphQLSample/src/interfaces/repositories"
	"github.com/gin-gonic/gin"
)

type HobbyController struct {
	Usecase *masters.HobbyUsecase
}

func NewHobbyController(db *db.Database) *HobbyController {
	return &HobbyController{
		Usecase: &masters.HobbyUsecase{
			HobbyRepository: &repositories.HobbyRepository{},
			DB:              db,
		},
	}
}
func (controller *HobbyController) GetHobbies(c *gin.Context) {
	output, err := controller.Usecase.GetHobbies()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, output)
}

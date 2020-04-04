package users

import (
	"github.com/GraphQLSampleGqlgen/src/db"
	"github.com/GraphQLSampleGqlgen/src/model"
	"github.com/GraphQLSampleGqlgen/src/usecases/repositories"
)

type UserUsecase struct {
	UserRepository repositories.UserRepository
	DB             *db.Database
}

func (usecase *UserUsecase) FindUser(input uint) (*model.User, error) {
	user, err := usecase.UserRepository.SelectByUserID(usecase.DB.MainDB.ReadReplica, input)
	if err != nil {
		return nil, err
	}

	todos := []*model.Todo{}

	for _, todo := range user.UserTodos {
		todos = append(todos, &model.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
		})
	}

	output := &model.User{
		ID:             user.ID,
		DisplayID:      user.DisplayID,
		IsUnsubscribed: user.IsUnsubscribed,
		Nickname:       user.UserDetail.Nickname,
		Birthday:       user.UserDetail.Birthday,
		Todos:          todos,
	}
	return output, nil
}

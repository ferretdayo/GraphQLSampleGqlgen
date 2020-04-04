package users

import (
	"errors"

	"github.com/GraphQLSampleGqlgen/src/db"
	"github.com/GraphQLSampleGqlgen/src/entities"
	"github.com/GraphQLSampleGqlgen/src/model"
	"github.com/GraphQLSampleGqlgen/src/usecases/repositories"
)

type UserTodoUsecase struct {
	UserTodoRepository repositories.UserTodoRepository
	DB                 *db.Database
}

func (usecase *UserTodoUsecase) FindUserTodo(input uint) ([]*model.Todo, error) {
	userTodos, err := usecase.UserTodoRepository.SelectByUserID(usecase.DB.MainDB.ReadReplica, input)
	if err != nil {
		return nil, errors.New("something wrong.")
	}
	var output []*model.Todo
	for _, userTodo := range userTodos {
		output = append(output, &model.Todo{
			ID:   userTodo.ID,
			Text: userTodo.Text,
			Done: userTodo.Done,
		})
	}
	return output, nil
}

func (usecase *UserTodoUsecase) CreateUserTodo(input model.NewTodo) (*model.Todo, error) {
	userTodo := &entities.UserTodo{
		UserID: input.UserID,
		Text:   input.Text,
	}
	if err := usecase.UserTodoRepository.Insert(usecase.DB.MainDB.Master, userTodo); err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:   userTodo.ID,
		Text: userTodo.Text,
		Done: userTodo.Done,
	}, nil
}

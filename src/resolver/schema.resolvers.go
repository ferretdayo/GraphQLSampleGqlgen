// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolver

import (
	"context"
	"errors"
	"fmt"

	"github.com/GraphQLSampleGqlgen/src/generated"
	"github.com/GraphQLSampleGqlgen/src/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context, userID uint) ([]*model.Todo, error) {
	userTodos, err := r.Resolver.UserTodoRepository.SelectByUserID(r.Resolver.DB.MainDB.ReadReplica, userID)
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
	fmt.Printf("todo: #%v", output)
	return output, nil
}

func (r *queryResolver) User(ctx context.Context, id uint) (*model.User, error) {
	user, err := r.Resolver.UserRepository.SelectByUserID(r.DB.MainDB.ReadReplica, id)
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

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

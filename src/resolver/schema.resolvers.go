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

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	userTodos, err := r.Resolver.UserTodoRepository.Select(r.Resolver.DB.MainDB.ReadReplica)
	if err != nil {
		return nil, errors.New("something wrong.")
	}
	var output []*model.Todo
	for _, userTodo := range userTodos {
		output = append(output, &model.Todo{
			ID:   userTodo.ID,
			Text: userTodo.Text,
			Done: userTodo.Done,
			User: &model.User{
				ID:             userTodo.User.ID,
				DisplayID:      userTodo.User.DisplayID,
				IsUnsubscribed: userTodo.User.IsUnsubscribed,
				Nickname:       userTodo.User.UserDetail.Nickname,
				Birthday:       userTodo.User.UserDetail.Birthday,
			},
		})
	}
	fmt.Printf("todo: #%v", output)
	return output, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.Resolver.UserRepository.Select(r.Resolver.DB.MainDB.ReadReplica)
	if err != nil {
		return nil, errors.New("something wrong.")
	}
	fmt.Printf("user: %v\n", users[0].UserTodos)
	output := []*model.User{}
	for _, user := range users {
		todos := []*model.Todo{}
		for _, userTodo := range user.UserTodos {
			todos = append(todos, &model.Todo{
				ID:   userTodo.ID,
				Text: userTodo.Text,
				Done: userTodo.Done,
			})
		}
		output = append(output, &model.User{
			ID:             user.ID,
			DisplayID:      user.DisplayID,
			IsUnsubscribed: user.IsUnsubscribed,
			Nickname:       user.UserDetail.Nickname,
			Birthday:       user.UserDetail.Birthday,
			Todos:          todos,
		})
	}

	fmt.Printf("output: #%v\n", output)
	return output, nil
}

func (r *queryResolver) User(ctx context.Context, id uint) (*model.User, error) {
	user, err := r.Resolver.UserRepository.SelectByUserID(r.DB.MainDB.ReadReplica, id)
	if err != nil {
		return nil, err
	}

	output := &model.User{
		ID:             user.ID,
		DisplayID:      user.DisplayID,
		IsUnsubscribed: user.IsUnsubscribed,
		Nickname:       user.UserDetail.Nickname,
		Birthday:       user.UserDetail.Birthday,
	}
	fmt.Printf("user: %v\n", output)
	return output, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

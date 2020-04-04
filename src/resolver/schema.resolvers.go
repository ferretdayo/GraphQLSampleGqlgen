// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolver

import (
	"context"

	"github.com/GraphQLSampleGqlgen/src/generated"
	"github.com/GraphQLSampleGqlgen/src/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	output, uerr := r.Resolver.UserTodoUsecase.CreateUserTodo(input)
	if uerr != nil {
		return nil, uerr
	}
	return output, nil
}

func (r *queryResolver) Todos(ctx context.Context, userID uint) ([]*model.Todo, error) {
	output, uerr := r.Resolver.UserTodoUsecase.FindUserTodo(userID)
	if uerr != nil {
		return nil, uerr
	}
	return output, nil
}

func (r *queryResolver) User(ctx context.Context, id uint) (*model.User, error) {
	output, uerr := r.Resolver.UserUsecase.FindUser(id)
	if uerr != nil {
		return nil, uerr
	}
	return output, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

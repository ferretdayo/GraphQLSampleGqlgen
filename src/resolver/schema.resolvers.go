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
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.Resolver.UserRepository.Select(r.Resolver.DB.MainDB.ReadReplica)
	if err != nil {
		return nil, errors.New("something wrong.")
	}

	var output []*model.User
	for _, user := range users {
		output = append(output, &model.User{
			ID:             string(user.ID),
			DisplayID:      user.DisplayID,
			IsUnsubscribed: user.IsUnsubscribed,
		})
	}
	return output, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

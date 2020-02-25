// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"

	generated1 "github.com/GraphQLSampleGqlgen/src/graph/generated"
	model1 "github.com/GraphQLSampleGqlgen/src/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model1.NewTodo) (*model1.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model1.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model1.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated1.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

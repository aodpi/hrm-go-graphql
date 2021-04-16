package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aodpi/hrm-go-graphql/v2/graph/generated"
	"github.com/aodpi/hrm-go-graphql/v2/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "1",
			Text: "Test text",
			Done: true,
			User: &model.User{
				ID:   "1",
				Name: "Test user",
			},
		},
	}, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	return &model.Todo{
		ID: "asdasdasd",
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

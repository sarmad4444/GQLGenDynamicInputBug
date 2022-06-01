package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"todo"
)

func (r *mutationsResolver) CreateTodo(ctx context.Context, input map[string]interface{}) (*todo.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// ERROR HERE
func (r *mutationsResolver) CreateTodoMany(ctx context.Context, inputs map[string]interface{}) ([]*todo.Todo, error) {
	// Incorrect Parameter Type Generated
	return nil, errors.New("Error here, expexted []map[string]interface{} but gqlgen generated map[string]interface")
}

func (r *queriesResolver) Todos(ctx context.Context) ([]*todo.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutations returns MutationsResolver implementation.
func (r *Resolver) Mutations() MutationsResolver { return &mutationsResolver{r} }

// Queries returns QueriesResolver implementation.
func (r *Resolver) Queries() QueriesResolver { return &queriesResolver{r} }

type mutationsResolver struct{ *Resolver }
type queriesResolver struct{ *Resolver }

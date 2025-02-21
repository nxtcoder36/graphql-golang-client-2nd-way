package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"github.com/nxtCoder36/graphql-golang-client-2nd-way/graph/model"
)

// AddTodoItem is the resolver for the addTodoItem field.
func (r *mutationResolver) AddTodoItem(ctx context.Context, title string) (*model.Todo, error) {
	todo := &model.Todo{
		ID:    len(r.TodoElements) + 1,
		Title: title,
	}
	r.TodoElements = append(r.TodoElements, todo)
	return todo, nil
	//panic(fmt.Errorf("not implemented: AddTodoItem - addTodoItem"))
}

// UpdateTodoItem is the resolver for the updateTodoItem field.
func (r *mutationResolver) UpdateTodoItem(ctx context.Context, id int, title *string, completed *bool) (*bool, error) {
	for i := range r.TodoElements {
		if r.TodoElements[i].ID == id {
			if title != nil {
				r.TodoElements[i].Title = *title
			}
			if completed != nil {
				r.TodoElements[i].Completed = *completed
			}
			result := true
			return &result, nil
		}
	}
	result := false
	return &result, nil
	//panic(fmt.Errorf("not implemented: UpdateTodoItem - updateTodoItem"))
}

// DeleteTodoItem is the resolver for the deleteTodoItem field.
func (r *mutationResolver) DeleteTodoItem(ctx context.Context, id int) (*bool, error) {
	for i := range r.TodoElements {
		if r.TodoElements[i].ID == id {
			r.TodoElements = append(r.TodoElements[:i], r.TodoElements[i+1:]...)
			result := true
			return &result, nil
		}
	}
	result := true
	return &result, nil
	//panic(fmt.Errorf("not implemented: DeleteTodoItem - deleteTodoItem"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	//panic(fmt.Errorf("not implemented: Todos - todos"))
	return r.TodoElements, nil

}

// TodoByID is the resolver for the todoByID field.
func (r *queryResolver) TodoByID(ctx context.Context, id int) (*model.Todo, error) {
	for i := range r.TodoElements {
		if r.TodoElements[i].ID == id {
			return r.TodoElements[i], nil
		}
	}
	return nil, nil
	//panic(fmt.Errorf("not implemented: TodoByID - todoByID"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

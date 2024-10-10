package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"project-test/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
// curl --location 'localhost:8080/query' \
// --header 'Content-Type: application/json' \
// --header 'Authorization: ••••••' \
//
//	--data '{
//	   "query": "mutation { createTodo(input: {userId: \"id1\", text: \"john.doe\"}) { id text done user { id name } } }"
//	}'
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", randNumber),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.Resolver.todo = append(r.Resolver.todo, todo)
	return todo, nil
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todo := []*model.Todo{
		{
			ID:   "ID1",
			Text: "A example response",
			Done: true,
			User: &model.User{
				ID:   "UserID1",
				Name: "Test",
			},
		},
	}

	return todo, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return obj.User, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *orderResolver) Bill(ctx context.Context, obj *model.Order) (*model.Bill, error) {
	panic(fmt.Errorf("not implemented: Bill - bill"))
}
func (r *Resolver) Order() OrderResolver { return &orderResolver{r} }
type orderResolver struct{ *Resolver }
*/

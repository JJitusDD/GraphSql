//go:generate go run github.com/99designs/gqlgen generate
package graph

import "project-test/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todo  []*model.Todo
	order []*model.Order
}

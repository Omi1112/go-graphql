package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/SeijiOmi/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver 解決
type Resolver struct {
	todos []*model.Todo
	users []*model.User
}

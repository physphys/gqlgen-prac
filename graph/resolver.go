package graph

import "github.com/physphys/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// in-memory database
type Resolver struct {
	todos []*model.Todo
}

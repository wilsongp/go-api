package graphql

import (
	"github.com/wilsongp/go-api/graphql/todos"

	"github.com/graphql-go/graphql"
)

// RootMutation is the base mutation for all graphql schema
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "RootMutation",
	Fields: todos.TodoMutations,
})

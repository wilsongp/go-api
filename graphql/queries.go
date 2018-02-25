package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/wilsongp/go-api/graphql/todos"
)

// RootQuery
// we just define a trivial example here, since root query is required.
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:   "RootQuery",
	Fields: todos.TodoQuery,
})

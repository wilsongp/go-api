package todos

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/wilsongp/go-api/graphql/common/repository"
)

//TodoQuery : base todo query
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var TodoQuery = graphql.Fields{

	/*
	   curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'
	*/
	"todo": &graphql.Field{
		Type:        TodoType,
		Description: "Get single todo",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			idQuery, isOK := params.Args["id"].(string)

			if !isOK {
				fmt.Println("invalid ID supplied")
			}

			_, database := repository.DialServer(repository.SERVER, repository.DBNAME)

			todo, _ := database.GetDocumentByID("todo", idQuery)

			return todo, nil
		},
	},

	/*
	   curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'
	*/
	"todoList": &graphql.Field{
		Type:        graphql.NewList(TodoType),
		Description: "List of todos",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			_, database := repository.DialServer(repository.SERVER, repository.DBNAME)

			todos, _ := database.GetDocuments("todo", p.Source)

			return todos, nil
		},
	},
}

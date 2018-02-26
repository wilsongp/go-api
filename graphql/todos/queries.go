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

			_, conn := repository.DialServer(repository.SERVER)

			var todo Todo
			// database.C(COLLECTION).FindId(idQuery).One(&todo)

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

			_, conn := repository.DialServer(repository.SERVER)

			var results []Todo
			if err := database.C(COLLECTION).Find(nil).Limit(100).All(&results); err != nil {
				fmt.Println("Failed to write get:", err)
				return results, err
			}

			return results, nil
		},
	},
}

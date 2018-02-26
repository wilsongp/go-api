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
				fmt.Println("invalid ID supplied: ", idQuery)
			}

			_, _, err := repository.DialServer(repository.SERVER)
			if err != nil {
				// do some awesome error handling
			}

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

			_, conn, err := repository.DialServer(repository.SERVER)
			if err != nil {
				fmt.Println("Error connecting to database: ", err)
			}

			cypher := `MATCH (todo:Todo) RETURN todo.id as id, todo.text as text, todo.done as done LIMIT {limit}`
			data, _, _, err := conn.QueryNeoAll(cypher, map[string]interface{}{"limit": 100})

			results := make([]Todo, len(data))
			for i, row := range data {
				results[i] = Todo{
					ID:   row[0].(string),
					Text: row[1].(string),
					Done: row[2].(bool),
				}
			}

			return results, err
		},
	},
}

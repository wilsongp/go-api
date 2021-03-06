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
				return nil, fmt.Errorf("Invalid idQuery: %s", idQuery)
			}

			_, conn, err := repository.DialServer(repository.SERVER)
			if err != nil {
				fmt.Println("Error connecting to database: ", err)
			}
			defer conn.Close()

			cypher := `MATCH (todo:Todo) WHERE todo.id = {id} RETURN todo.id as id, todo.text as text, todo.done as done LIMIT {limit}`
			cypherParams := map[string]interface{}{
				"id":    idQuery,
				"limit": 1,
			}

			var result Todo
			data, _, _, err := conn.QueryNeoAll(cypher, cypherParams)
			if len(data) == 1 {
				result = Todo{
					ID:   data[0][0].(string),
					Text: data[0][1].(string),
					Done: data[0][2].(bool),
				}
			}

			return result, err
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
			defer conn.Close()

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

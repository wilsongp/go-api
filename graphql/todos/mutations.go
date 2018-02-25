package todos

import (
	"log"

	"github.com/wilsongp/go-api/graphql/common/repository"

	"github.com/graphql-go/graphql"
)

// TodoMutations : Base mutation for todos schema
var TodoMutations = graphql.Fields{
	/*
	   curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'
	*/
	"createTodo": &graphql.Field{
		Type:        TodoType, // the return type for this field
		Description: "Create new todo",
		Args: graphql.FieldConfigArgument{
			"text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			// marshall and cast the argument value
			text, _ := params.Args["text"].(string)
			isDone, _ := params.Args["isDone"].(bool)

			_, database := repository.DialServer(repository.SERVER, repository.DBNAME)

			todo := Todo{
				text: text,
				done: isDone,
			}

			if _, err := database.InsertDocument(todo, "todo"); err != nil {
				log.Fatal(err)
				return nil, err
			}

			return todo, nil
		},
	},
	/*
	   curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"a",done:true){id,text,done}}'
	*/
	"updateTodo": &graphql.Field{
		Type:        TodoType, // the return type for this field
		Description: "Update existing todo, mark it done or not done",
		Args: graphql.FieldConfigArgument{
			"done": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// marshall and cast the argument value
			done, _ := params.Args["done"].(bool)
			id, _ := params.Args["id"].(string)

			_, database := repository.DialServer(repository.SERVER, repository.DBNAME)

			var updatedTodo Todo
			updatedTodo.done = done

			if _, err := database.UpdateDocumentByID(id, updatedTodo, "todo"); err != nil {
				log.Fatal(err)
				return nil, err
			}

			return updatedTodo, nil
		},
	},
}

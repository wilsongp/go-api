package todos

import (
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

			_, connection := repository.DialServer(repository.SERVER, repository.DBNAME)

			todo := Todo{
				Text: text,
				Done: isDone,
			}

			// if err := database.C(COLLECTION).Insert(todo); err != nil {
			// 	log.Fatal(err)
			// 	return nil, err
			// }

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

			_, connection := repository.DialServer(repository.SERVER, repository.DBNAME)

			var updatedTodo Todo
			updatedTodo.Done = done

			// if err := database.C(COLLECTION).UpdateId(id, updatedTodo); err != nil {
			// 	log.Fatal(err)
			// 	return updatedTodo, err
			// }

			return updatedTodo, nil
		},
	},
}
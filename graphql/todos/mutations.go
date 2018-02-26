package todos

import (
	"github.com/nu7hatch/gouuid"
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
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// marshall and cast the argument value
			text, _ := p.Args["text"].(string)
			isDone, _ := p.Args["isDone"].(bool)
			newGUID, _ := uuid.NewV4()

			params := map[string]interface{}{
				"text": text,
				"done": isDone,
				"id":   newGUID.String(),
			}

			_, conn, dialerr := repository.DialServer(repository.SERVER)
			if dialerr != nil {
				// do some awesome error handling
			}

			cypher := `CREATE (n:Todo { id: {id}, text: {text}, done: {done} }) RETURN n`
			_, err := conn.ExecNeo(cypher, params)

			result := Todo{
				ID:   newGUID.String(),
				Text: text,
				Done: isDone,
			}

			return result, err
		},
	},
	/*
	   curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"a",done:true){id,text,done}}'
	*/
}

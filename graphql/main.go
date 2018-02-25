package graphql

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/wilsongp/go-api/routing"

	"github.com/graphql-go/graphql"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

//Routes for graphql
var Routes = routing.Routes{
	{
		Name:    "GraphQL",
		Method:  "GET",
		Pattern: "/graphql",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			result := executeQuery(r.URL.Query().Get("query"), schema)
			json.NewEncoder(w).Encode(result)
			return
		},
	},
}

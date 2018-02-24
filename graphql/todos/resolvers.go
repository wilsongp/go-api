package todos

import (
  "log"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SERVER the DB server
const SERVER = "localhost:32768"

// DBNAME the name of the DB instance
const DBNAME = "homepage"

// DOCNAME the name of the document
const DOCNAME = "shortcuts"

//CreateTodoResolver adds a ToDo to the database
func CreateTodoResolver(params graphql.ResolveParams) (interface{}, error) {

  // marshall and cast the argument value
  text, _ := params.Args["text"].(string)
  isDone, _ := params.Args["isDone"].(bool)

  session, err := mgo.Dial(SERVER)
	defer session.Close()

  todo := Todo {
    ID: bson.NewObjectId(),
    text: text,
    done: isDone,
  }
	session.DB(DBNAME).C(DOCNAME).Insert(todo)

	if err != nil {
		log.Fatal(err)
	}

  // return the new Todo object that we supposedly save to DB
  // Note here that
  // - we are returning a `Todo` struct instance here
  // - we previously specified the return Type to be `todos.TodoType`
  // - `Todo` struct maps to `todos.TodoType`, as defined in `todos.TodoType` ObjectConfig`
  return todo, nil
}
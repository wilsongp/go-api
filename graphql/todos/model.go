package todos

import (
	"gopkg.in/mgo.v2/bson"
)

// COLLECTION : The name of the collection in the db
const COLLECTION = "todos"

//Todo document model
type Todo struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Text string        `json:"text"`
	Done bool          `json:"done"`
}

//Todos is an array of Todo
type Todos []Todo

package todos

import (
	"gopkg.in/mgo.v2/bson"
)

//Todo document model
type Todo struct {
	ID   bson.ObjectId `bson:"_id"`
	text string        `json:"text"`
	done bool          `json:"isDone"`
}

//Todos is an array of Todo
type Todos []Todo

package shortcut

import "gopkg.in/mgo.v2/bson"

//Shortcut represents link
type Shortcut struct {
	id   bson.ObjectId `bson:"_id"`
	href string        `json:"title"`
	text string        `json:"artist"`
	sort int32         `json:"year"`
}

//Shortcuts is an array of Album
type Shortcuts []Shortcut

package shortcut

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:32768"

// DBNAME the name of the DB instance
const DBNAME = "homepage"

// DOCNAME the name of the document
const DOCNAME = "shortcuts"

// GetShortcuts returns the list of Shortcuts
func (r Repository) GetShortcuts() Shortcuts {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Shortcuts{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddShortcut inserts an Shortcut in the DB
func (r Repository) AddShortcut(shortcut Shortcut) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	shortcut.id = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(shortcut)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateShortcut updates an Shortcut in the DB
func (r Repository) UpdateShortcut(shortcut Shortcut) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	shortcut.id = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).UpdateId(shortcut.id, shortcut)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteShortcut deletes an Shortcut
func (r Repository) DeleteShortcut(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "404"
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "500"
	}

	// Write status
	return "200"
}

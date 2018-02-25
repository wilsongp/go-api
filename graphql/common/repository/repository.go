package repository

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

// SERVER the DB server
const SERVER = "mongodb://192.168.99.100:32768/"

// DBNAME the name of the DB instance
const DBNAME = "homepage"

// DOCNAME the name of the document
const DOCNAME = "shortcuts"

//Repository ...
type Repository struct{}

// Database : Extends mgo.Database with generic CRUD methods
type Database struct {
	*mgo.Database
}

// Session : Extends mgo.Session
type Session struct {
	*mgo.Session
}

var _sessions = make(map[string]Session)

// DialServer : Opens session on supplied server
func DialServer(connection string, database string) (Session, Database) {

	if _, ok := _sessions[connection]; !ok {
		session, err := mgo.Dial(connection)
		if err != nil {
			fmt.Printf("\nFailed to establish connection to Mongo server '%s':\n'%s'\n\n", connection, err)
		}

		_sessions[connection] = Session{session}
		fmt.Println("Connected to Mongo server at: ", connection)
	}

	db := _sessions[connection].DB(database)
	return _sessions[connection], Database{db}
}

// GetDocuments : Executes a query against the database
func (db Database) GetDocuments(doctype string, query interface{}) ([]interface{}, error) {
	c := db.C(doctype)

	var results []interface{}
	if err := c.Find(query).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
		return results, err
	}

	return results, nil
}

// GetDocumentByID : Gets a Mongo doc given doctype and an id string
func (db Database) GetDocumentByID(doctype string, id string) (interface{}, error) {
	c := db.C(doctype)

	var result interface{}
	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write results:", err)
		return result, err
	}

	return result, nil
}

// InsertDocument : Add document to database
func (db Database) InsertDocument(newDoc interface{}, doctype string) (bool, error) {
	var reflected = newDoc
	if err := db.C(doctype).Insert(reflected); err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

// UpdateDocumentByID : Update document to database
func (db Database) UpdateDocumentByID(id string, updated interface{}, doctype string) (bool, error) {

	if err := db.C(doctype).UpdateId(id, updated); err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

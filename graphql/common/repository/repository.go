package repository

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// SERVER the DB server
const SERVER = "bolt://localhost:32769"

// DBNAME the name of the DB instance
const DBNAME = "homepage"

// DOCNAME the name of the document
const DOCNAME = "shortcuts"

//Repository ...
type Repository struct{}

// Connection : Extends mgo.Database with generic CRUD methods
type Connection struct {
	bolt.Conn
}

// Driver : Extends mgo.Driver
type Driver struct {
	bolt.Driver
}

var _driverPools = make(map[string]Driver)

//MAX_CONNECTIONS : Max number of connections to pool
const MAX_CONNECTIONS = 5

// DialServer : Opens session on supplied server
func DialServer(connection string) (Driver, Connection, error) {

	if _, ok := _driverPools[connection]; !ok {
		driver := bolt.NewDriver()

		_driverPools[connection] = Driver{driver}
		fmt.Println("Connected to Neo4j server at: ", connection)
	}

	db, err := _driverPools[connection].OpenNeo(connection)
	if err != nil {
		return Driver{}, Connection{}, err
	}

	return _driverPools[connection], Connection{db}, nil
}

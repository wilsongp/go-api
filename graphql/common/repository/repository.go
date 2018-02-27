package repository

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// SERVER the DB server
const SERVER = "bolt://neo4j:test@localhost:32769"

// DBNAME the name of the DB instance
const DBNAME = "homepage"

// DOCNAME the name of the document
const DOCNAME = "shortcuts"

//MAX_CONNECTIONS : Max number of connections to pool
const MAX_CONNECTIONS = 5

//Repository ...
type Repository struct{}

// Connection : Extends mgo.Database with generic CRUD methods
type Connection struct {
	bolt.Conn
}

// Driver : Extends mgo.Driver
type Driver struct {
	bolt.DriverPool
}

var _driverPools = make(map[string]Driver)

// DialServer : Opens session on supplied server
func DialServer(connection string) (Driver, Connection, error) {

	if _, ok := _driverPools[connection]; !ok {
		driver, err := bolt.NewDriverPool(connection, MAX_CONNECTIONS)
		if err != nil {
			return Driver{}, Connection{}, err
		}

		_driverPools[connection] = Driver{driver}
		fmt.Println("Connected to Neo4j server...")
	}

	conn, err := _driverPools[connection].OpenPool()
	if err != nil {
		return Driver{}, Connection{}, err
	}

	return _driverPools[connection], Connection{conn}, nil
}

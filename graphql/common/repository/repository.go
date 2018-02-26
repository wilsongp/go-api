package repository

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// SERVER the DB server
const SERVER = "mongodb://192.168.99.100:32768/"

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
	bolt.DriverPool
}

var _driverPools = make(map[string]Driver)

//MAX_CONNECTIONS : Max number of connections to pool
const MAX_CONNECTIONS = 5

// DialServer : Opens session on supplied server
func DialServer(connection string) (Driver, Connection) {

	if _, ok := _driverPools[connection]; !ok {
		driver, err := bolt.NewDriverPool(connection, MAX_CONNECTIONS)
		if err != nil {
			panic(err)
		}

		_driverPools[connection] = Driver{driver}
		fmt.Println("Connected to Neo4j server at: ", connection)
	}

	db, poolerr := _driverPools[connection].OpenPool()
	if poolerr != nil {
		panic(poolerr)
	}
	return _driverPools[connection], Connection{db}
}

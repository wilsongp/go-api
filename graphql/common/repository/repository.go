package repository

import (
	"fmt"
	"os"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// SERVER the DB server
var SERVER = "bolt://" + os.Getenv("NEO4J_USER") + ":" + os.Getenv("NEO4J_PASSWORD") + "@" + os.Getenv("NEO4J_HOST") + ":" + os.Getenv("NEO4J_PORT")

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
func DialServer(host string) (Driver, Connection, error) {

	if _, ok := _driverPools[host]; !ok {
		driver, err := bolt.NewDriverPool(host, MAX_CONNECTIONS)
		if err != nil {
			return Driver{}, Connection{}, err
		}

		_driverPools[host] = Driver{driver}
		fmt.Println("Connected to Neo4j server...")
	}

	conn, err := _driverPools[host].OpenPool()
	if err != nil {
		return Driver{}, Connection{}, err
	}

	return _driverPools[host], Connection{conn}, nil
}

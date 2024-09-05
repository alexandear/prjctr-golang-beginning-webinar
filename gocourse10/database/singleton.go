package database

import (
	"log"
	"sync"
)

type Connection interface {
	Open() error
	Close() error
}

type Executor interface {
	Execute(cmd string) error
}

type MedicalDatabase struct {
	connection Connection
}

func (c *MedicalDatabase) SetConnection(connection Connection) {
	log.Printf("Setting connection: %#+v\n", connection)
	c.connection = connection
}

func (c *MedicalDatabase) Connect() error {
	return c.connection.Open()
}

func (c *MedicalDatabase) Disconnect() error {
	return c.connection.Close()
}

func (c *MedicalDatabase) Execute(cmd string) error {
	log.Printf("Query %v executed\n", cmd)
	return nil
}

var (
	onceDatabase sync.Once
	database     Database
)

func Instance() Database {
	onceDatabase.Do(func() {
		database = &MedicalDatabase{
			connection: &SQLConnection{},
		}

		log.Printf("Database instance created: %#+v", database)
	})
	return database
}

type Database interface {
	SetConnection(connection Connection)
	Connect() error
	Disconnect() error
	Execute(cmd string) error
}

func SetInstance(db Database) {
	database = db
	log.Printf("Database instance set: %#+v", database)
}

package database

import (
	"log"
	"math/rand/v2"
	"time"
)

type LogDecorator struct {
	wrapped Database
	log     *log.Logger
}

func NewLogDecorator(wrapped Database, log *log.Logger) Database {
	return &LogDecorator{
		wrapped: wrapped,
		log:     log,
	}
}

func (d *LogDecorator) Execute(cmd string) error {
	start := time.Now()
	defer func() { d.log.Printf("Database query took: %v\n", time.Since(start)) }()
	randDelay()
	return d.wrapped.Execute(cmd)
}

func (d *LogDecorator) Connect() error {
	start := time.Now()
	defer func() { d.log.Printf("Database connection took: %v\n", time.Since(start)) }()
	randDelay()
	return d.wrapped.Connect()
}

func (d *LogDecorator) Disconnect() error {
	start := time.Now()
	defer func() { d.log.Printf("Database disconnection took: %v\n", time.Since(start)) }()
	randDelay()
	return d.wrapped.Disconnect()
}

func (d *LogDecorator) SetConnection(connection Connection) {
	d.wrapped.SetConnection(connection)
}

func randDelay() {
	time.Sleep(time.Duration(100+rand.N(200)) * time.Millisecond)
}

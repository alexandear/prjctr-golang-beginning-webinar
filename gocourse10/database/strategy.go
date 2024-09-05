package database

import "log"

type SQLConnection struct{}

func (s *SQLConnection) Open() error {
	log.Println("SQL Database Connection Opened")
	return nil
}

func (s *SQLConnection) Close() error {
	log.Println("SQL Database Connection Closed")
	return nil
}

var _ Connection = &SQLConnection{}

type NoSQLConnection struct{}

func (n *NoSQLConnection) Open() error {
	log.Println("NoSQL Database Connection Opened")
	return nil
}

func (n *NoSQLConnection) Close() error {
	log.Println("NoSQL Database Connection Closed")
	return nil
}

var _ Connection = &NoSQLConnection{}

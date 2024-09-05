package main

import (
	"log"

	"github.com/alexandear/prjctr-golang-beginning-webinar/gocourse10/database"
)

func main() {
	db := database.Instance()
	err := db.Connect()
	if err != nil {
		log.Panic(err)
	}

	err = db.Execute("Create")
	if err != nil {
		log.Panic(err)
	}
	err = db.Execute("Update")
	if err != nil {
		log.Panic(err)
	}
	err = db.Execute("Delete")
	if err != nil {
		log.Panic(err)
	}

	err = db.Disconnect()
	if err != nil {
		log.Printf("Disconnecting: %v\n", err)
	}

	sqlConn := &database.NoSQLConnection{}
	db.SetConnection(sqlConn)
	err = db.Disconnect()
	if err != nil {
		log.Printf("Disconnecting: %v\n", err)
	}

	logDB := database.NewLogDecorator(db, log.Default())
	database.SetInstance(logDB)
	db = database.Instance()

	err = db.Execute("Create")
	if err != nil {
		log.Panic(err)
	}
	err = db.Execute("Update")
	if err != nil {
		log.Panic(err)
	}

	err = db.Disconnect()
	if err != nil {
		log.Printf("Disconnecting: %v\n", err)
	}
}

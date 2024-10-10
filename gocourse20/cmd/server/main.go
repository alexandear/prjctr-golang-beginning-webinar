package main

import (
	"github.com/gorilla/mux"
	"gocourse20/core/service/patients"
	"gocourse20/interfaces/rest/handlers"
	middlewares2 "gocourse20/interfaces/rest/middlewares"
	"gocourse20/pkg/telementry/meter"
	"log"
	"net/http"
)

func main() {
	meter.MustInit(`0.0.0.0:4343`, true)

	patientsService := patients.NewService()

	// start the rest server
	restHandler := handlers.NewPatients(patientsService)

	router := mux.NewRouter()

	router.HandleFunc("/patients", restHandler.AddPatient).Methods("POST")
	router.HandleFunc("/patients/{id}", restHandler.GetPatient).Methods("GET")
	router.HandleFunc("/patients/{id}", restHandler.UpdatePatient).Methods("PUT")

	handler := middlewares2.Meter(middlewares2.Last(router))

	server := &http.Server{Addr: ":8080", Handler: handler}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return
}

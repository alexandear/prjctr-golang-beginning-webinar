package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"gocourse20/internal/rest/handler"
	"gocourse20/internal/rest/middleware"
	"gocourse20/internal/service/patient"
	"gocourse20/internal/telemetry/meter"
)

func main() {
	meter.MustInit("localhost:4343")

	patientService := patient.NewService()

	restHandler := handler.NewPatient(patientService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /patients", restHandler.AddPatient)
	mux.HandleFunc("GET /patients/{id}", restHandler.GetPatient)
	mux.HandleFunc("PUT /patients/{id}", restHandler.UpdatePatient)

	handler := middleware.Meter(middleware.Last(mux))

	server := &http.Server{
		Addr:              "localhost:8080",
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("RESTful server running on http://%s\n", server.Addr)

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failed to serve: %v", err)
	}
}

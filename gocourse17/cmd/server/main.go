package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"prjctr.com/gocourse17/internal/grpc/adapter"
	"prjctr.com/gocourse17/internal/grpc/grpcapi"
	"prjctr.com/gocourse17/internal/rest/handler"
	"prjctr.com/gocourse17/internal/service/patient"
)

func main() {
	patientService := patient.NewService()

	go func() {
		server := adapter.NewPatient(patientService)

		lis, err := net.Listen("tcp", "localhost:50051")
		if err != nil {
			log.Fatalf("Failed to listen gRPC: %v", err)
		}
		s := grpc.NewServer()
		grpcapi.RegisterPatientServiceServer(s, server)

		log.Printf("gRPC server serving on %s", lis.Addr().String())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	handler := handler.NewPatient(patientService)

	router := http.NewServeMux()

	router.Handle("POST /patients", http.HandlerFunc(handler.AddPatient))
	router.HandleFunc("GET /patients/{id}", handler.GetPatient)
	router.HandleFunc("PUT /patients/{id}", handler.UpdatePatient)

	server := &http.Server{
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to listen HTTP: %v", err)
	}
	log.Printf("HTTP server serving on %s", lis.Addr().String())
	if err := server.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		log.Panicf("failed to serve: %v", err)
	}
}

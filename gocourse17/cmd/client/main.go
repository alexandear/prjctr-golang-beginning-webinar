package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"prjctr.com/gocourse17/internal/grpc/grpcapi"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := grpcapi.NewPatientServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddPatient(ctx, &grpcapi.AddPatientRequest{Patient: &grpcapi.Patient{Id: "1", Name: "John Doe", Age: "30", Diagnosis: "Diagnosis"}})
	if err != nil {
		log.Panicf("Failed to add patient: %v", err)
	}
	log.Printf("Server response: %s", r.GetMessage())

	r2, err := c.GetPatient(ctx, &grpcapi.GetPatientRequest{Id: "1"})
	if err != nil {
		log.Panicf("Failed to get patient: %v", err)
	}
	log.Printf("Patient: %v", r2.GetPatient())

	r3, err := c.UpdatePatient(ctx, &grpcapi.UpdatePatientRequest{
		Patient: &grpcapi.Patient{Id: "1", Name: "John Doe", Age: "31", Diagnosis: "Updated diagnosis"},
	})
	if err != nil {
		log.Panicf("Failed to update patient: %v", err)
	}
	log.Printf("Server response: %s", r3.GetMessage())
}

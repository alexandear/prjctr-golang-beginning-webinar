package adapter

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"prjctr.com/gocourse17/internal/grpc/grpcapi"
	"prjctr.com/gocourse17/internal/service/patient"
)

func NewPatient(s *patient.Service) *Patient {
	return &Patient{service: s}
}

type Patient struct {
	service *patient.Service
	grpcapi.UnimplementedPatientServiceServer
}

func (s *Patient) AddPatient(ctx context.Context, in *grpcapi.AddPatientRequest) (*grpcapi.AddPatientResponse, error) {
	inPatient := in.GetPatient()
	patientEntity := patient.Patient{
		ID:        inPatient.GetId(),
		Name:      inPatient.GetName(),
		Age:       inPatient.GetAge(),
		Diagnosis: inPatient.GetDiagnosis(),
	}

	if _, err := s.service.AddPatient(ctx, &patientEntity); err != nil {
		return nil, fmt.Errorf("failed to add patient: %w", err)
	}

	return &grpcapi.AddPatientResponse{Message: "Patient added successfully"}, nil
}

func (s *Patient) GetPatient(ctx context.Context, in *grpcapi.GetPatientRequest) (*grpcapi.GetPatientResponse, error) {
	res, err := s.service.GetPatient(ctx, in.GetId())
	switch {
	case err == nil:
		return &grpcapi.GetPatientResponse{Patient: &grpcapi.Patient{
			Id:        res.ID,
			Name:      res.Name,
			Age:       res.Age,
			Diagnosis: res.Diagnosis,
		}}, nil
	case errors.Is(err, patient.ErrNotFound):
		return nil, status.Errorf(codes.NotFound, "patient not found: %v", err)
	default:
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}
}

func (s *Patient) UpdatePatient(ctx context.Context, in *grpcapi.UpdatePatientRequest) (*grpcapi.UpdatePatientResponse, error) {
	inPatient := in.GetPatient()
	patientEntity := patient.Patient{
		ID:        inPatient.GetId(),
		Name:      inPatient.GetName(),
		Age:       inPatient.GetAge(),
		Diagnosis: inPatient.GetDiagnosis(),
	}
	if _, err := s.service.UpdatePatient(ctx, in.GetPatient().GetId(), &patientEntity); err != nil {
		return nil, fmt.Errorf("failed to update patient: %w", err)
	}
	return &grpcapi.UpdatePatientResponse{Message: "Patient updated successfully"}, nil
}

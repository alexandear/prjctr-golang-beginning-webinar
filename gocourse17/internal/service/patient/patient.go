package patient

import (
	"context"
	"errors"
	"log"
)

var ErrNotFound = errors.New("patient not found")

type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       string `json:"age"`
	Diagnosis string `json:"diagnosis"`
}

func NewService() *Service {
	return &Service{}
}

type Service struct {
}

func (s *Service) AddPatient(ctx context.Context, in *Patient) (*Patient, error) {
	log.Println("Adding patient", in)
	return in, nil
}

func (s *Service) GetPatient(ctx context.Context, id string) (*Patient, error) {
	log.Println("Getting patient", id)
	return &Patient{ID: id}, nil
}

func (s *Service) UpdatePatient(ctx context.Context, id string, in *Patient) (*Patient, error) {
	log.Println("Updating patient", id, in)
	return in, nil
}

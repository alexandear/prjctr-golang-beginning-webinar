package patient

import (
	"context"
	"log"
	"math/rand/v2"
	"time"
)

type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       string `json:"age"`
	Diagnosis string `json:"diagnosis"`
}

func NewService() *Service {
	return &Service{}
}

type Service struct{}

func (s *Service) AddPatient(ctx context.Context, in *Patient) (*Patient, error) {
	simulateLatency()
	log.Println("Adding patient", in)
	return in, nil
}

func (s *Service) GetPatient(ctx context.Context, id string) (*Patient, error) {
	simulateLatency()
	log.Println("Getting patient", id)
	return &Patient{ID: id}, nil
}

func (s *Service) UpdatePatient(ctx context.Context, id string, in *Patient) (*Patient, error) {
	simulateLatency()
	log.Println("Updating patient", id, in)
	return in, nil
}

func simulateLatency() {
	time.Sleep(time.Duration(rand.N(2000)+200) * time.Millisecond)
}

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gocourse20/core/service/patients"
	"log"
	"net/http"
	"time"
)

func NewPatients(s *patients.Service) *Patients {
	return &Patients{s}
}

type Patients struct {
	service *patients.Service
}

// Додавання пацієнта
func (p *Patients) AddPatient(w http.ResponseWriter, r *http.Request) {
	var patient *patients.Patient

	time.Sleep(1 * time.Second)

	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		log.Fatal()
	}

	res, err := p.service.AddPatient(r.Context(), patient)
	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		log.Fatal()
	}

	_ = json.NewEncoder(w).Encode(res)
}

// Отримання пацієнта
func (p *Patients) GetPatient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	time.Sleep(1 * time.Second)

	res, err := p.service.GetPatient(r.Context(), id)
	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		log.Fatal()
	}

	_ = json.NewEncoder(w).Encode(res)
}

// Оновлення пацієнта
func (p *Patients) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	time.Sleep(1 * time.Second)

	var updatedPatient patients.Patient
	if err := json.NewDecoder(r.Body).Decode(&updatedPatient); err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		log.Fatal()
	}

	res, err := p.service.UpdatePatient(r.Context(), id, &updatedPatient)
	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		log.Fatal()
	}

	_ = json.NewEncoder(w).Encode(res)
}

package handler

import (
	"encoding/json"
	"net/http"

	"prjctr.com/gocourse17/internal/service/patient"
)

func NewPatient(s *patient.Service) *Patient {
	return &Patient{s}
}

type Patient struct {
	service *patient.Service
}

func (p *Patient) AddPatient(w http.ResponseWriter, r *http.Request) {
	var patient *patient.Patient

	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := p.service.AddPatient(r.Context(), patient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(res)
}

func (p *Patient) GetPatient(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	res, err := p.service.GetPatient(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(res)
}

func (p *Patient) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var updatedPatient patient.Patient
	if err := json.NewDecoder(r.Body).Decode(&updatedPatient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := p.service.UpdatePatient(r.Context(), id, &updatedPatient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(res)
}

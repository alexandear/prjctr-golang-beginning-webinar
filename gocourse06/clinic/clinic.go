package clinic

import "sync"

type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BloodType string `json:"blood_type"`
}

type Clinic struct {
	patients map[string]Patient
	mu       sync.Mutex
}

func NewClinic() *Clinic {
	return &Clinic{
		patients: make(map[string]Patient),
	}
}

func (c *Clinic) AddPatient(p Patient) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.patients[p.ID] = p
}

func (c *Clinic) Patient(id string) (Patient, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	p, exists := c.patients[id]
	return p, exists
}

func (c *Clinic) Patients() map[string]Patient {
	return c.patients
}

func (c *Clinic) ForcedLock() {
}

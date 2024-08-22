package clinic

type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BloodType string `json:"blood_type"`
}

type Clinic struct {
	patients map[string]Patient
}

func NewClinic() *Clinic {
	return &Clinic{
		patients: make(map[string]Patient),
	}
}

func (c *Clinic) AddPatient(p Patient) {
	c.patients[p.ID] = p
}

func (c *Clinic) Patient(id string) (Patient, bool) {
	p, exists := c.patients[id]
	return p, exists
}

func (c *Clinic) Patients() map[string]Patient {
	return c.patients
}

func (c *Clinic) ForcedLock() {
}

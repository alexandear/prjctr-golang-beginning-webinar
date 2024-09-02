package clinic

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Clinic struct {
	mu       sync.RWMutex
	patients map[string]Patient

	dataChan chan string
	done     atomic.Bool
}

func NewClinic() *Clinic {
	return &Clinic{
		patients: make(map[string]Patient),
		dataChan: make(chan string),
	}
}

func (c *Clinic) Chan() <-chan string {
	return c.dataChan
}

func (c *Clinic) AddPatient(p Patient) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.patients[p.ID] = p
}

func (c *Clinic) AddPatientWhileCtx(ctx context.Context, p <-chan Patient) {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Adding patient cancelled\n")
			return
		case patient := <-p:
			c.mu.Lock()
			c.patients[patient.ID] = patient
			c.mu.Unlock()
			fmt.Println("Patient added:", patient.ID)
		case <-t.C:
			fmt.Println("Just tick")
		}
	}
}

func (c *Clinic) DeletePatient(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.patients, id)
}

func (c *Clinic) Stop(done chan<- struct{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for id := range c.patients {
		delete(c.patients, id)
		time.Sleep(time.Second / 10)
	}

	c.done.Store(true)
	close(c.dataChan)

	done <- struct{}{}
}

func (c *Clinic) ProcessData(patientID int, p Patient, gg <-chan struct{}, wayChan chan<- chan string) {
	defer func() { <-gg }()

	if c.done.Load() {
		fmt.Printf("Patient %d won't be processed\n", patientID)
		return
	}

	var controlChan chan string
	if patientID%10 == 0 {
		controlChan = make(chan string)
		wayChan <- controlChan
	}

	var patientData string
	select {
	case way := <-controlChan:
		patientData = fmt.Sprintf("Patient %d processed in Special way: %s", patientID, way)
	case <-time.After(time.Second):
		patientData = fmt.Sprintf("Patient %d processed", patientID)
	}

	c.dataChan <- patientData
	p.Data = patientData

	c.AddPatient(p)
}

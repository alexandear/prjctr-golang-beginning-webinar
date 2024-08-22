package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/alexandear/prjctr-golang-beginning-webinar/gocourse06/clinic"
)

func main() {
	c := clinic.NewClinic()
	patients := []clinic.Patient{
		{"1", "John Doe", 30, "A+"},
		{"2", "Jane Smith", 25, "O-"},
		{"3", "Will Smith", 50, "C-"},
	}

	var wg sync.WaitGroup
	for _, p := range patients {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.AddPatient(p)
			fmt.Printf("Patient added: %s\n", p.Name)
		}()
	}
	wg.Wait()
	fmt.Printf("%d patients added to the clinic.\n", len(c.Patients()))

	bloodTypes := []string{"A+", "O-", "B+", "AB+"}
	ch := make(chan string, len(bloodTypes))

	for i, bt := range bloodTypes {
		go processBloodType(bt, ch, i)
	}
	for range bloodTypes {
		fmt.Println(<-ch)
	}

	addPatientChan := make(chan clinic.Patient)
	getPatientChan := make(chan string)
	quitChan := make(chan bool)

	go func() {
		for {
			select {
			case p := <-addPatientChan:
				c.AddPatient(p)
			case id := <-getPatientChan:
				patient, exists := c.Patient(id)
				if exists {
					fmt.Printf("Patient retrieved: %s, Age: %d, BloodType: %s\n", patient.Name, patient.Age, patient.BloodType)
				} else {
					fmt.Printf("Patient not found with ID: %s", id)
					return
				}
			case <-quitChan:
				fmt.Println("Clinic closed.")
				return
			}
		}
	}()

	addPatientChan <- clinic.Patient{"1", "Mike Doe", 35, "A+"}
	addPatientChan <- clinic.Patient{"2", "Mika Smith", 20, "O-"}
	// c.ForcedLock()
	addPatientChan <- clinic.Patient{"3", "Mika Smith", 33, "AO-"}

	getPatientChan <- "1"
	getPatientChan <- "3"

	quitChan <- true
}

func processBloodType(bloodType string, ch chan<- string, sec int) {
	time.Sleep(time.Duration(sec+1) * time.Second)
	ch <- fmt.Sprintf("Processed blood type: %s", bloodType)
}

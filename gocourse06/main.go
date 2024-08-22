package main

import (
	"fmt"
	"sync"

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
}

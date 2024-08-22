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
}

func processBloodType(bloodType string, ch chan<- string, sec int) {
	time.Sleep(time.Duration(sec+1) * time.Second)
	ch <- fmt.Sprintf("Processed blood type: %s", bloodType)
}

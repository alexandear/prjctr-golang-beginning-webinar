package main

import (
	"fmt"

	"github.com/alexandear/prjctr-golang-beginning-webinar/gocourse07/clinic"
)

const totalPatients = 100
const maxGoroutines = 3

func main() {
	patients := make([]clinic.Patient, totalPatients)
	cl := clinic.NewClinic()

	for i := range totalPatients {
		patients[i] = clinic.GenerateRandomPatient()
	}

	// Горутина для системи сповіщення лікарів
	go func() {
		for p := range cl.Chan() {
			fmt.Printf("Doctor was notified about patient: %s\n", p)
		}
	}()

	wayChan := make(chan chan string)
	go func() {
		for where := range wayChan {
			where <- clinic.GenerateRandomString(10)
		}
	}()

	// Старт горутин для обробки даних пацієнтів
	closeClinic := make(chan struct{})
	gGuard := make(chan struct{}, maxGoroutines)
	go func() {
		for i, patient := range patients {
			select {
			case <-closeClinic:
				fmt.Println("Clinic closed")
				return
			default:
				gGuard <- struct{}{} // would block if gGuard channel is already filled
				go cl.ProcessData(i, patient, gGuard, wayChan)
			}
		}
	}()

	// time.Sleep(5 * time.Second)
	close(closeClinic)
	done := make(chan struct{}, 1)
	go cl.Stop(done)
	<-done

	fmt.Println("All patients processed for today")
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://localhost:8080/patients"

type Patient struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Age       string `json:"age"`
	Diagnosis string `json:"diagnosis"`
}

func addPatient(patient Patient) {
	payload, err := json.Marshal(patient)
	if err != nil {
		fmt.Println("Error marshaling patient:", err)
		return
	}

	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Add Patient Response:", string(body))
}

func getPatient(id string) {
	url := fmt.Sprintf("%s/%s", baseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Get Patient Response:", string(body))
}

func updatePatient(id string, patient Patient) {
	url := fmt.Sprintf("%s/%s", baseURL, id)
	payload, err := json.Marshal(patient)
	if err != nil {
		fmt.Println("Error marshaling patient:", err)
		return
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Update Patient Response:", string(body))
}

func main() {
	newPatient := Patient{
		Name:      "John Doe",
		Age:       "30",
		Diagnosis: "Flu",
	}
	addPatient(newPatient)

	// Припускаючи, що ID нового пацієнта відомий та є "1", для прикладу
	getPatient("1")

	updatedPatient := Patient{
		Name:      "John Doe Updated",
		Age:       "31",
		Diagnosis: "Cold",
	}
	updatePatient("1", updatedPatient)
}

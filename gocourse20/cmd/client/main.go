package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type patient struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Age       string `json:"age"`
	Diagnosis string `json:"diagnosis"`
}

type patientClient struct {
	client  *http.Client
	baseURL string
}

func (c *patientClient) addPatient(ctx context.Context, patient patient) (string, error) {
	payload, err := json.Marshal(patient)
	if err != nil {
		return "", fmt.Errorf("marshaling patient: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response: %w", err)
	}

	return string(body), nil
}

func (c *patientClient) getPatient(ctx context.Context, id string) (string, error) {
	url := c.baseURL + "/" + id
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response: %w", err)
	}

	return string(body), nil
}

func (c *patientClient) updatePatient(ctx context.Context, id string, patient patient) (string, error) {
	payload, err := json.Marshal(patient)
	if err != nil {
		return "", fmt.Errorf("marshaling patient: %w", err)
	}

	url := c.baseURL + "/" + id
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response: %w", err)
	}

	return string(body), nil
}

func main() {
	pc := &patientClient{
		client:  http.DefaultClient,
		baseURL: "http://localhost:8080/patients",
	}

	addResp, err := pc.addPatient(context.Background(), patient{
		Name:      "John Doe",
		Age:       "30",
		Diagnosis: "Flu",
	})
	if err != nil {
		log.Fatalln("Error adding patient:", err)
	}
	log.Println("Add patient response:", addResp)

	getResp, err := pc.getPatient(context.Background(), "1")
	if err != nil {
		log.Fatalln("Error getting patient:", err)
	}
	log.Println("Get patient response:", getResp)

	updateResp, err := pc.updatePatient(context.Background(), "1", patient{
		Name:      "John Doe Updated",
		Age:       "31",
		Diagnosis: "Cold",
	})
	if err != nil {
		log.Fatalln("Error updating patient:", err)
	}
	log.Println("Update patient response:", updateResp)
}

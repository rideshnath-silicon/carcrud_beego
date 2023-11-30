package test

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestCarContrller(t *testing.T) {
	t.Run("Insert New Car", func(t *testing.T) {
		endPoint := "http://localhost:8080/v1/car/create"

		var jsonStr = []byte(`{
				"car_name" : "Thar",
				"model": "4*4",
				"modified_by": "mahindara",
				"type":"sedan"
			}`)

		req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}
		responseString := string(body)

		log.Print(responseString)
	})

	t.Run("Update Car", func(t *testing.T) {
		endPoint := "http://localhost:8080/v1/car/update"

		var jsonStr = []byte(`{
				"car_id":1,
				"car_name" : "Thar",
				"model": "4*4",
				"modified_by": "mahindara",
				"type":"sedan"
			}`)
		req, err := http.NewRequest("PUT", endPoint, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}
		responseString := string(body)
		log.Print(responseString)
	})
	t.Run("Getcars", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "http://localhost:8080/v1/car/cars", nil)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		body, _ := io.ReadAll(resp.Body)
		log.Print(string(body))
	})
	t.Run("Search car", func(t *testing.T) {
		endPoint := "http://localhost:8080/v1/car/search"

		var jsonStr = []byte(`{
				"search":"thar"
			}`)
		req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}
		responseString := string(body)
		log.Print(responseString)
	})
}

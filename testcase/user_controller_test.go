package test

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
)

// func Init() {
// 	orm.RegisterDriver("postgres", orm.DRPostgres)
// 	orm.RegisterDataBase("default", "postgres", "user=postgres password=root dbname=mydb sslmode=disable")
// 	orm.RunSyncdb("default", false, true)
// 	flag.Parse()
// }

func TestUserControllers(t *testing.T) {
	t.Run("register", func(t *testing.T) {
		endPoint := "http://localhost:8080/v1/user/register"

		var jsonStr = []byte(`{
		"first_name" : "Dwarkesh",
		"last_name" : "patel",
		"email" : "dwarkeshpatel.siliconithub@gmail.com",
		"country" : "India",
		"role" : "Developer",
		"age" : 30,
		"phone_number":"1234567890",
		"password" : "123456"
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
	t.Run("tesst All user Controller", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "http://localhost:8080/v1/user/users", nil)
		client := &http.Client{}
		resp, _ := client.Do(req)
		body, _ := io.ReadAll(resp.Body)
		log.Print(string(body))
	})
	t.Run("Update user", func(t *testing.T) {
		endPoint := "http://localhost:8080/v1/user/update"

		var jsonStr = []byte(`{
			"user_id":9,
			"first_name" : "Demo",
			"last_name" : "test",
			"email" : "demo@gmal.com",
			"country" : "India",
			"role" : "demo",
			"age" : 12,
			"phone_number" : "1234567890"
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
	t.Run("Verify email ", func(t *testing.T) {
		endPoint := "http://localhost:8080/v1/user/verify_email"

		var jsonStr = []byte(`{
			"username":"rideshnath.siliconithub@gmail.com"
		}`)

		req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}
		responseString := string(body)
		log.Print(responseString)
	})
	t.Run("Verify email  otp ", func(t *testing.T) {
		endPoint := "http://localhost:8080/v1/user/verify_email_otp"

		var jsonStr = []byte(`{
			"username": "6264736064", 
			"otp":"6773"
		}`)

		req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}
		responseString := string(body)
		log.Print(responseString)
	})
	t.Run("Forgot  Password ", func(t *testing.T) {
		endPoint := "http://localhost:8080/v1/user/forgot_pass"

		var jsonStr = []byte(`{
			"username":"rideshnath.siliconithub@gmail.com"
		}`)

		req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}
		responseString := string(body)
		log.Print(responseString)
	})
	t.Run("Reset  Password ", func(t *testing.T) {
		endPoint := "http://localhost:8080/v1/user/reset_pass_otp"

		var jsonStr = []byte(`{
			"email" : "rideshnath.siliconithub@gmail.com",
			"otp":"0703",
		"new_password":"123456"
		}`)

		req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
			return
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

func TestLogin(t *testing.T) {
	t.Run("login", func(t *testing.T) {
		endPoint := "http://localhost:8080/v1/login"

		var jsonStr = []byte(`{
			"username" : "6264736064",
			"password": "123456"
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

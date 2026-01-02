package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"auth/models"
)

func TestRegister(t *testing.T) {
	// Reset users for test
	models.Users = make(map[string]models.User)

	payload := []byte(`{"username":"testuser","password":"password123","permissions":1}`) // 1 = Read
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check if user exists in model
	if _, exists := models.Users["testuser"]; !exists {
		t.Errorf("user was not created")
	}
}

func TestLogin(t *testing.T) {
	// Reset users
	models.Users = make(map[string]models.User)

	// Register a user first
	regPayload := []byte(`{"username":"loginuser","password":"password123","permissions":1}`)
	regReq, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(regPayload))
	regRr := httptest.NewRecorder()
	http.HandlerFunc(Register).ServeHTTP(regRr, regReq)

	// Now try to login
	payload := []byte(`{"username":"loginuser","password":"password123"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("failed to parse response: %v", err)
	}

	if _, ok := response["token"]; !ok {
		t.Errorf("token not returned in response")
	}
}

func TestLoginInvalidPassword(t *testing.T) {
	models.Users = make(map[string]models.User)

	// Register
	regPayload := []byte(`{"username":"wrongpassuser","password":"password123","permissions":1}`)
	regReq, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(regPayload))
	http.HandlerFunc(Register).ServeHTTP(httptest.NewRecorder(), regReq)

	// Login with wrong password
	payload := []byte(`{"username":"wrongpassuser","password":"wrongpassword"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}

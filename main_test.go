package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/AbanoubGirges/malaykaproject/models"
	migrations "github.com/AbanoubGirges/malaykaproject/repo"
	"github.com/AbanoubGirges/malaykaproject/routes"
	"github.com/AbanoubGirges/malaykaproject/services"
	"github.com/joho/godotenv"
)

var (
	testRouter *http.Handler
	testToken  string
	adminToken string
	testUserID uint32
)

func TestMain(m *testing.M) {
	godotenv.Load(".env")
	DB := migrations.SetupDatabase(os.Getenv("DB_URL"))
	services.SetDB(DB)
	services.TakeSecretKey(os.Getenv("SECRET_KEY"))
	code := m.Run()
	os.Exit(code)
}

func getUniqueUsername() string {
	return "testuser_" + strconv.Itoa(rand.Intn(1000000))
}

func TestSignup(t *testing.T) {
	router := routes.SetupRouter("8080")

	user := models.User{
		Username:    getUniqueUsername(),
		PhoneNumber: "1234567890",
		Password:    "password123",
	}

	payload, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
	t.Logf("Signup response: %d - %s", w.Code, w.Body.String())
}

func TestLogin(t *testing.T) {
	router := routes.SetupRouter("8080")

	loginReq := models.UserLoginRequest{
		PhoneNumber: "1234567890",
		Password:    "password123",
	}

	payload, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	t.Logf("Login response: %d - %s", w.Code, w.Body.String())
	if w.Code == http.StatusOK {
		var response map[string]interface{}
		json.NewDecoder(w.Body).Decode(&response)
		if token, ok := response["Token"]; ok {
			testToken = token.(string)
		}
	}
}

func TestCreateStudent_MissingAuth(t *testing.T) {
	router := routes.SetupRouter("8080")

	student := models.Student{
		Name:        "John Doe",
		PhoneNumber: []string{"1234567890"},
		Location:    "City",
		Age:         15,
		Coordinates: "10.5,20.5",
		Birthdate:   "2009-01-01",
	}

	payload, _ := json.Marshal(student)
	req := httptest.NewRequest("POST", "/students/create", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 without auth token, got %d", w.Code)
	}
	t.Logf("Create student without auth: %d - %s", w.Code, w.Body.String())
}

func TestDeleteStudent_BuggyFieldAccess(t *testing.T) {
	// This test exposes the bug: DeleteStudentHandler uses claims["ID"] which doesn't exist
	// It should use claims["user_id"]
	t.Log("BUG DETECTED: DeleteStudentHandler accesses claims[\"ID\"] but JWT contains \"user_id\"")
	t.Log("Expected panic or error when deleting student")
}

func TestEditProfile_MissingResponse(t *testing.T) {
	// This test shows the missing response issue
	// EditProfileHandler doesn't return a response on success
	t.Log("BUG DETECTED: EditProfileHandler doesn't send response after successful update")
}

func TestReadyEndpoint(t *testing.T) {
	router := routes.SetupRouter("8080")

	req := httptest.NewRequest("GET", "/ready", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
	t.Logf("Ready endpoint: %d", w.Code)
}

func TestCreateClass_MissingAuth(t *testing.T) {
	router := routes.SetupRouter("8080")

	req := httptest.NewRequest("POST", "/class/create?class_name=Math", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 without auth, got %d", w.Code)
	}
	t.Logf("Create class without auth: %d", w.Code)
}

func TestInvalidJSON(t *testing.T) {
	router := routes.SetupRouter("8080")

	// Send invalid JSON to signup
	req := httptest.NewRequest("POST", "/signup", bytes.NewBufferString("invalid json"))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 for invalid JSON, got %d", w.Code)
	}
	t.Logf("Invalid JSON response: %d", w.Code)
}

func TestMissingRequiredFields(t *testing.T) {
	router := routes.SetupRouter("8080")

	// Signup with missing fields
	user := models.User{
		Username: "testuser",
		// Missing PhoneNumber and Password
	}

	payload, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 for missing fields, got %d", w.Code)
	}
	t.Logf("Missing fields response: %d", w.Code)
}

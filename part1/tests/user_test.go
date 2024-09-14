package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CeyhunBoran/shaffra-casestudy/internal/config"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/handlers"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/models"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/repositories"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/services"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/utils"
)

var (
	testUser    models.User
	testRepo    repositories.UserRepository
	testService services.UserService
	testHandler handlers.UserHandler
)

func setupTests() {
	config.InitConfigTest()

	conf := config.Conf
	db, err := utils.NewDB(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		conf.DbHost, conf.DbUser, conf.DbPass, conf.DbTestName, conf.DbTestPort, conf.DbSsl))
	if err != nil {
		fmt.Printf("Failed to create database connection: %v", err)
		return
	}
	db.Conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	testRepo = *repositories.NewUserRepository(db)
	testService = *services.NewUserService(testRepo)
	testHandler = *handlers.NewUserHandler(testService)
}

func teardownTests(failed bool) {
	// Clean up any resources created during tests
}

func TestMain(m *testing.M) {
	setupTests()
	code := m.Run()
	teardownTests(code == 0)
}

func TestCreateUser(t *testing.T) {
	// Setup
	setupTests()

	// Arrange
	testUser = models.User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	// Act
	reqBody, _ := json.Marshal(testUser)
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	testHandler.CreateUser(response, req)

	// Assert
	body, _ := ioutil.ReadAll(response.Body)
	var result models.User
	json.Unmarshal(body, &result)

	t.Logf("Created User: %+v", result)
	t.Log("Status:", response.Code)

	if response.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", response.Code)
	}

	if result.ID == "" {
		t.Error("User ID is empty")
	}

	if result.Name != testUser.Name {
		t.Errorf("Expected Name to be '%s', got '%s'", testUser.Name, result.Name)
	}

	if result.Email != testUser.Email {
		t.Errorf("Expected Email to be '%s', got '%s'", testUser.Email, result.Email)
	}

	if result.Age != testUser.Age {
		t.Errorf("Expected Age to be %d, got %d", testUser.Age, result.Age)
	}
}

func TestGetUser(t *testing.T) {
	// Setup
	setupTests()

	// Arrange
	testUser := models.User{
		Name:  "Jane Doe",
		Email: "jane@example.com",
		Age:   25,
	}

	// Create the user
	createdUser, err := testService.CreateUser(testUser)
	if err != nil {
		t.Fatal(err)
	}

	// Construct the request URL with the actual user ID
	userID := createdUser.ID
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/users/%s", userID), nil)
	response := httptest.NewRecorder()
	testHandler.GetUser(response, req)

	// Assert
	body, _ := ioutil.ReadAll(response.Body)
	var result models.User
	json.Unmarshal(body, &result)

	t.Logf("Retrieved User: %+v", result)
	t.Log("Status:", response.Code)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}

	if result.ID == "" {
		t.Error("User ID is empty")
	}

	if result.Name != testUser.Name {
		t.Errorf("Expected Name to be '%s', got '%s'", testUser.Name, result.Name)
	}

	if result.Email != testUser.Email {
		t.Errorf("Expected Email to be '%s', got '%s'", testUser.Email, result.Email)
	}

	if result.Age != testUser.Age {
		t.Errorf("Expected Age to be %d, got %d", testUser.Age, result.Age)
	}
}

func TestUpdateUser(t *testing.T) {
	// Setup
	setupTests()

	// Arrange
	testUser = models.User{
		Name:  "Updated John",
		Email: "updated@example.com",
		Age:   31,
	}

	// Act
	_, err := testService.CreateUser(models.User{Name: "Original John"})
	if err != nil {
		t.Fatal(err)
	}

	reqBody, _ := json.Marshal(testUser)
	req, _ := http.NewRequest("PUT", "/api/users/original-id", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	testHandler.UpdateUser(response, req)

	// Assert
	body, _ := ioutil.ReadAll(response.Body)
	var result models.User
	json.Unmarshal(body, &result)

	t.Logf("Updated User: %+v", result)
	t.Log("Status:", response.Code)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}

	if result.ID == "" {
		t.Error("User ID is empty")
	}

	if result.Name != testUser.Name {
		t.Errorf("Expected Name to be '%s', got '%s'", testUser.Name, result.Name)
	}

	if result.Email != testUser.Email {
		t.Errorf("Expected Email to be '%s', got '%s'", testUser.Email, result.Email)
	}

	if result.Age != testUser.Age {
		t.Errorf("Expected Age to be %d, got %d", testUser.Age, result.Age)
	}
}

func TestDeleteUser(t *testing.T) {
	// Setup
	setupTests()

	// Arrange
	testUser = models.User{
		Name:  "Deletable User",
		Email: "deletable@example.com",
		Age:   28,
	}

	// Act
	userToDelete, err := testService.CreateUser(testUser)
	if err != nil {
		t.Fatal(err)
	}
	userID := userToDelete.ID
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%s", userID), nil)
	response := httptest.NewRecorder()
	testHandler.DeleteUser(response, req)

	// Assert
	t.Log("Status:", response.Code)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code 204, got %d", response.Code)
	}

	// Try to get deleted user
	getReq, _ := http.NewRequest("GET", "/api/users/deletable-id", nil)
	getResponse := httptest.NewRecorder()
	testHandler.GetUser(getResponse, getReq)

	if getResponse.Code != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %d", getResponse.Code)
	}
}

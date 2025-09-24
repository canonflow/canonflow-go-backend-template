package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"canonflow-golang-backend-template/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	viperConfig := viper.New()

	wd, _ := os.Getwd()
	projectRoot := filepath.Join(wd, "..") // move up one dir from /test
	viperConfig.SetConfigFile(filepath.Join(projectRoot, ".env"))

	err := viperConfig.ReadInConfig()

	assert.Nil(t, err)
}

func InitServer() *gin.Engine {
	viperConfig := viper.New()

	wd, _ := os.Getwd()
	projectRoot := filepath.Join(wd, "..") // move up one dir from /test
	viperConfig.SetConfigFile(filepath.Join(projectRoot, ".env"))

	err := viperConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}

	log := config.NewLogrus(viperConfig)
	validate := config.NewValidator()
	app := config.NewGin(viperConfig, log)
	db := config.NewDatabase(viperConfig, log)

	// Bootstrap all configs
	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	return app
}

func TestSignUpSuccess(t *testing.T) {
	app := InitServer()

	// Request Body
	body := map[string]string{
		"username": "username_test",
		"password": "password",
	}

	// Encode Body
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)

	request := httptest.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	// Record
	recorder := httptest.NewRecorder()
	app.ServeHTTP(recorder, request)

	// Response
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	/*
		=== RUN   TestSignUpSuccess
		--- PASS: TestSignUpSuccess (0.07s)
		PASS
	*/
}

func TestSignUpFailed(t *testing.T) {
	app := InitServer()

	// Request Body
	body := map[string]string{
		"username": "username_test",
		"password": "password",
	}

	// Encode Body
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)

	request := httptest.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	// Record
	recorder := httptest.NewRecorder()
	app.ServeHTTP(recorder, request)

	// Response
	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	/*
		=== RUN   TestSignUpFailed
		--- PASS: TestSignUpFailed (0.01s)
		PASS
	*/
}

func TestLoginSuccess(t *testing.T) {
	app := InitServer()

	// Request Body
	body := map[string]string{
		"username": "username_test",
		"password": "password",
	}

	// Encode Body
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)

	request := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	// Record
	recorder := httptest.NewRecorder()
	app.ServeHTTP(recorder, request)

	// Response
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	/*
		=== RUN   TestLoginSuccess
		--- PASS: TestLoginSuccess (0.07s)
		PASS
	*/
}

func TestLoginNotFound(t *testing.T) {
	app := InitServer()

	// Request Body
	body := map[string]string{
		"username": "username_test_not_found",
		"password": "password",
	}

	// Encode Body
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)

	request := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	// Record
	recorder := httptest.NewRecorder()
	app.ServeHTTP(recorder, request)

	// Response
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	/*
		--- PASS: TestLoginNotFound (0.01s)
		=== RUN   TestLoginNotFound
		PASS
	*/
}

func TestLoginWrongPassword(t *testing.T) {
	app := InitServer()

	// Request Body
	body := map[string]string{
		"username": "username_test",
		"password": "wrong_password",
	}

	// Encode Body
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)

	request := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	// Record
	recorder := httptest.NewRecorder()
	app.ServeHTTP(recorder, request)

	// Response
	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	/*
		=== RUN   TestLoginWrongPassword
		--- PASS: TestLoginWrongPassword (0.07s)
		PASS
	*/
}

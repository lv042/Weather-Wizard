package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllWeatherData(t *testing.T) {
	// Initialize a Fiber application
	app := NewFiberApp()
	app.InitFiber()

	// Create a test HTTP request
	req := httptest.NewRequest(http.MethodGet, "/api/weather", nil)

	// Set the response for the test request
	resp, err := app.fiberApp.Test(req)
	if err != nil {
		t.Fatalf("Failed to run the test: %v", err)
	}

	// Check if the response status code is 200 OK
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
}

package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/covalenthq/covalent-api-sdk-go/utils"
)

func TestExponentialBackoff_BackOff(t *testing.T) {
	// Create a test server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	// Create a new ExponentialBackoff instance for testing
	backoff := utils.NewExponentialBackoff("API_KEY", true, 3)

	// Call the BackOff method with the test server URL
	response, err := backoff.BackOff(server.URL)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if response != nil {
		t.Errorf("Expected nil response, got %v", response)
	}
}

func TestExponentialBackoff_BackOffSuccess(t *testing.T) {
	// Create a test server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	// Create a new ExponentialBackoff instance for testing
	backoff := utils.NewExponentialBackoff("API_KEY", true, 3)

	// Call the BackOff method with the test server URL
	response, err := backoff.BackOff(server.URL)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
}

func TestExponentialBackoff_BackOffMaxRetriesExceeded(t *testing.T) {
	// Create a test server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	// Create a new ExponentialBackoff instance for testing
	backoff := utils.NewExponentialBackoff("API_KEY", true, 5)

	// Call the BackOff method with the test server URL
	_, err := backoff.BackOff(server.URL)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "max retries exceeded: 5" {
		t.Errorf("Expected error message 'max retries exceeded: 5', got '%s'", err.Error())
	}
}

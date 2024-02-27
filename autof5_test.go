package autof5

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAutoF5(t *testing.T) {
	respBody := "<html><body>Hello, World!</body></html>"
	// Create a mock HTTP handler for testing
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(respBody))
	})

	// Create a mock HTTP request
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	// Create a mock HTTP response recorder
	rec := httptest.NewRecorder()

	// Call the AutoF5 function with the mock handler and delay
	AutoF5Handler := AutoF5(mockHandler)

	// Serve the request using the AutoF5 handler
	AutoF5Handler.ServeHTTP(rec, req)

	// Check if the response body contains the AutoF5 script
	expectedBody := "fetch(\"/_autoF5_wait\", { mode: 'no-cors' }"

	if !strings.Contains(rec.Body.String(), expectedBody) {
		t.Errorf("AutoF5 script not found in response body")
	}

	// Check if the response code is 200 OK
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

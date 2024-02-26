package autof5

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLivereload(t *testing.T) {
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

	// Call the Livereload function with the mock handler and delay
	livereloadHandler := Livereload(mockHandler)

	// Serve the request using the livereload handler
	livereloadHandler.ServeHTTP(rec, req)

	// Check if the response body contains the livereload script
	expectedBody := "fetch(\"/_autoF5_wait\", { mode: 'no-cors' }"

	if !strings.Contains(rec.Body.String(), expectedBody) {
		t.Errorf("Livereload script not found in response body")
	}

	// Check if the response code is 200 OK
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

package cmd

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStatusCommand(t *testing.T) {
	// 1. Create a fake HTTP server that mimics the OpenAPI backend
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/status" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"version": "1.0.0", "status": "ok"}`))
			return
		}
		
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	// 2. Set the overridden URL flag or env so the client hits the test server
	// We'll update statusCmd temporarily to have an initial state if we had a flag
	// For this test, we can capture the standard output using bytes buffer
	
	// Create a buffer to capture CLI output
	b := new(bytes.Buffer)
	statusCmd.SetOut(b)
	statusCmd.SetErr(b)
	// Set the package-level override URL to our test server
	overrideServerURL = server.URL
	
	err := statusCmd.Execute()
	
	// We expect success now since it connects to the test server
	if err != nil {
		t.Fatalf("Expected command to execute cleanly, got error: %v", err)
	}

	out := b.String()
	if !strings.Contains(out, "Calling /status endpoint...") {
		t.Errorf("Expected output to contain 'Calling /status endpoint...', got: %s", out)
	}
}

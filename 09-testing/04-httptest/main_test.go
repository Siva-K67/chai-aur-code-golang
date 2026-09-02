package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	// httptest.NewRequest builds a FAKE request - no real network call happens
	req := httptest.NewRequest(http.MethodGet, "/status", nil)

	//httptest.NewRecorder() — creates a fake http.ResponseWriter that simply records whatever gets written to it (status code, headers, body),
	// instead of sending it over a real network connection
	rec := httptest.NewRecorder()

	// call the handler DIRECTLY - just a normal function call, no server involved
	statusHandler(rec, req)

	// check the status code
	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	// decode the JSON body that was written
	var response StatusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Status != "ok" {
		t.Errorf("expected status 'ok', got %q", response.Status)
	}
}

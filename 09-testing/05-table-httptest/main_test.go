package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// use go test -v to get ll the subtest output instead of the regular go test

func TestCoursesHandler(t *testing.T) {
	// the table - each row is one scenario to test
	testCases := []struct {
		name           string // describes this specific case
		method         string
		body           string // request body, "" for none
		expectedStatus int
	}{
		{
			name:           "GET all courses",
			method:         http.MethodGet,
			body:           "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "POST valid course",
			method:         http.MethodPost,
			body:           `{"name": "New Course", "price": 799}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "POST invalid JSON",
			method:         http.MethodPost,
			body:           `{invalid json`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "DELETE not allowed",
			method:         http.MethodDelete,
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range testCases {
		// t.Run creates a named subtest - shows up individually in test output
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/courses", strings.NewReader(tc.body))
			rec := httptest.NewRecorder()

			coursesHandler(rec, req)

			if rec.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, rec.Code)
			}
		})
	}
}

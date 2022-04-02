package controller_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmetcanaydemir/go-rest/pkg/api/controller"
)

func Test_GetInMemoryController(t *testing.T) {
	var tests = []struct {
		name           string
		query          string
		expectedStatus int
	}{
		{"successful request", "?key=test-key", http.StatusOK},
		{"key not found request", "?key=not-found-key", http.StatusNotFound},
		{"key empty request", "?key=", http.StatusBadRequest},
		{"wrong field request", "?mykey=test-key", http.StatusBadRequest},
		{"missing field request", "", http.StatusBadRequest},
	}
	inMemoryController := controller.InMemoryController{}
	controller.InMemoryDB.Store("test-key", "test-value")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/in-memory"+tt.query, nil)
			rr := httptest.NewRecorder()
			inMemoryController.InMemoryHandler(rr, req)

			gotStatus := rr.Code
			if gotStatus != tt.expectedStatus {
				t.Errorf("GetInMemoryController(%s) got %v, want %v", tt.query, gotStatus, tt.expectedStatus)
			}
		})
	}
}

func Test_PostInMemoryController(t *testing.T) {
	var tests = []struct {
		name           string
		bodyjson       string
		expectedStatus int
	}{
		{"successful request", `{"key":"test-key","value":"test-value"}`, http.StatusOK},
		{"key missing request", `{"value":"test-value"}`, http.StatusBadRequest},
		{"value missing request", `{"key":"test-key"}`, http.StatusBadRequest},
		{"key empty request", `{"key":","value":"test-value"}`, http.StatusBadRequest},
		{"wrong body request", `{wrongbody}`, http.StatusBadRequest},
	}
	inMemoryController := controller.InMemoryController{}
	controller.InMemoryDB.Store("test-key", "test-value")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/in-memory", bytes.NewBuffer([]byte(tt.bodyjson)))
			rr := httptest.NewRecorder()
			inMemoryController.InMemoryHandler(rr, req)

			gotStatus := rr.Code
			if gotStatus != tt.expectedStatus {
				t.Errorf("PostInMemoryController(%s) got %v, want %v", tt.name, gotStatus, tt.expectedStatus)
			}
		})
	}
}

func Test_NotAllowedInMemoryController(t *testing.T) {
	inMemoryController := controller.InMemoryController{}
	req, _ := http.NewRequest("PUT", "/in-memory", nil)
	rr := httptest.NewRecorder()
	inMemoryController.InMemoryHandler(rr, req)

	gotStatus := rr.Code
	if gotStatus != http.StatusMethodNotAllowed {
		t.Errorf("InMemoryController(%s) got %v, want %v", "method not allowed request", gotStatus, http.StatusMethodNotAllowed)
	}

}

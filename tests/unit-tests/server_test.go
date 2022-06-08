package tests

import (
	"algorithm-1/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: Write unit tests for generate schedule & check schedule
// when implementation is done.

func TestRootRoute(t *testing.T) {
	// Setup
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	// Act
	server.Root(response, request)

	// Assert
	actual_response := response.Body.String()
	expected_response := "Algorithm 1 REST server is alive!"

	if actual_response != expected_response {
		t.Errorf("got %q, want %q", actual_response, expected_response)
	}
}

package tests

import (
	"algorithm-1/server"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestCheckSchdulePass(t *testing.T) {
	// Setup
	jsonFile, err := os.Open("../data/working-schedule-test.json")
	if err != nil {
		t.Error("File not found")
	}

	request, _ := http.NewRequest(http.MethodPost, "/CheckSchedule", jsonFile)
	response := httptest.NewRecorder()

	// Act
	server.CheckSchedule(response, request)

	// Assert
	actual_response := response.Body.String()
	expected_response := "Schedule given is valid"

	if actual_response != expected_response {
		t.Errorf("got %q, want %q", actual_response, expected_response)
	}
}

func TestCheckScheduleProfFail(t *testing.T) {
	// Setup
	jsonFile, err := os.Open("../data/prof-violation-schedule-test.json")
	if err != nil {
		t.Error("File not found")
	}

	request, _ := http.NewRequest(http.MethodPost, "/CheckSchedule", jsonFile)
	response := httptest.NewRecorder()

	// Act
	server.CheckSchedule(response, request)

	// Assert
	actual_response := response.Body.String()
	expected_response := "error: Bird, Bill teaching another Fall course at MTh1300,   Schedule given has some violations that should be resolved"

	if actual_response != expected_response {
		t.Errorf("got %q, want %q", actual_response, expected_response)
	}
}

func TestCheckScheduleTimeslotFail(t *testing.T) {
	// Setup
	jsonFile, err := os.Open("../data/timeslot-violation-schedule-test.json")
	if err != nil {
		t.Error("File not found")
	}

	request, _ := http.NewRequest(http.MethodPost, "/CheckSchedule", jsonFile)
	response := httptest.NewRecorder()

	// Act
	server.CheckSchedule(response, request)

	// Assert
	actual_response := response.Body.String()
	expected_response := "error: SENG 275 is scheduled at same time as another required course SENG310 in Fall term,   Schedule given has some violations that should be resolved"

	if actual_response != expected_response {
		t.Errorf("got %q, want %q", actual_response, expected_response)
	}
}

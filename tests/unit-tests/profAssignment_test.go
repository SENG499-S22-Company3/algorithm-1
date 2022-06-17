package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"io/ioutil"
	"log"
	"testing"
)

func TestProfAssignment(t *testing.T) {
	jsonData, err := ioutil.ReadFile("../data/input-test.json")
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
	
	input, err := structs.ParseInput(jsonData)
	if err != nil {
		t.Error("Input parsing failed with error: ", err.Error())
	}

	if input.HistoricData.FallCourses == nil {
		t.Error("Input failed to be parsed")
	}

	testScheduleCourse := scheduling.AssignCourseProf(input.HistoricData.FallCourses, input.CoursesToSchedule.FallCourses, input.Professors)

	if testScheduleCourse[0].Prof.DisplayName == "" {
		t.Error("Professors not assigned to courses")
	}
}
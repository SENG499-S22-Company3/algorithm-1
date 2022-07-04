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
        log.Fatal("Error when opening input-test.json file: ", err)
    }
	
	input, err := structs.ParseInput(jsonData)
	if err != nil {
		t.Error("Input parsing failed with error: ", err.Error())
	}

	if input.HardScheduled.FallCourses == nil {
		t.Error("Input failed to be parsed: fall historical courses should not be null")
	}

	testScheduleCourse := scheduling.AssignCourseProf(input.HardScheduled.FallCourses, input.CoursesToSchedule.FallCourses, input.Professors)
	prefsMap, _ := scheduling.MapPreferences(input.Professors)

	for _,c := range testScheduleCourse {
		if c.Prof.DisplayName == "" {
			t.Error("Professors not assigned to course")
		}
		
		if val, pass := prefsMap[c.Prof.DisplayName][c.Subject+c.CourseNumber]; !pass && c.Prof.DisplayName != "TBD" {
			t.Error(c.Prof.DisplayName, "cannot teach this course since they have no (", val, ") preference.")
		}
	}
}
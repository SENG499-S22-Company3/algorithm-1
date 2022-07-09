package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"io/ioutil"
	"log"
	"testing"
)

func TestExtraDataScheduleAssignment(t *testing.T) {

	jsonData, err := ioutil.ReadFile("../data/extra-data-input-test.json")
	if err != nil {
		log.Fatal("Error when opening extra-data-input-test.json file: ", err)
	}

	input, err := structs.ParseInput(jsonData)
	if err != nil {
		t.Error("Input parsing failed with error: ", err.Error())
	}

	if input.HardScheduled.SpringCourses == nil {
		t.Error("Input failed to be parsed: fall hard schedules courses should not be null")
	}

	testStreamtype, err := scheduling.BaseTimeslotMaps(input.HardScheduled.FallCourses)
	if err != nil {
		t.Error(err)
	}

	testSchedule, _, err := scheduling.AddCoursesToStreamMaps(scheduling.Split(input.CoursesToSchedule.FallCourses), testStreamtype)
	if err != nil {
		t.Error(err)
	}

	testSchedule = scheduling.AssignCourseProf(input.HardScheduled.FallCourses, testSchedule, input.Professors)
	err = scheduling.ScheduleConstraintsCheck("Fall", testSchedule, input.Professors)
	if err != nil {
		t.Error(err)
	}

	err = scheduling.ScheduleConstraintsCheck("Fall", testSchedule, input.Professors)
	if err != nil {
		t.Error(err)
	}

	testSchedule = append(testSchedule, input.HardScheduled.FallCourses...)

	_, err = scheduling.BaseTimeslotMaps(testSchedule)
	if err != nil {
		t.Error(err)
	}
}
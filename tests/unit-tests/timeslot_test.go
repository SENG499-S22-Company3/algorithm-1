package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"io/ioutil"
	"os"
	"testing"
)

func TestBaseTimeslotMap(t *testing.T) {
	// Test data
	testAssignment := structs.Assignment{
		StartDate: "Sep 05, 2018",
		EndDate:   "Dec 05, 2018",
		BeginTime: "1130",
		EndTime:   "1250",
		HoursWeek: 3,
		Sunday:    false,
		Monday:    true,
		Tuesday:   false,
		Wednesday: false,
		Thursday:  true,
		Friday:    false,
		Saturday:  false,
	}

	testProf := structs.Professor{
		DisplayName: "JohnSmith",
	}

	testCourse := structs.Course{
		Assignment:     testAssignment,
		Prof:           testProf,
		CourseNumber:   "101",
		Subject:        "CHEM",
		SequenceNumber: "A01",
		StreamSequence: "2A",
		CourseTitle:    "Properties of Materials",
	}

	testSchedule := structs.Schedule{
		FallCourses:   []structs.Course{testCourse},
		SpringCourses: []structs.Course{},
		SummerCourses: []structs.Course{},
	}

	result, err := scheduling.BaseTimeslotMaps(testSchedule.FallCourses)

	if err != nil {
		t.Error(err)
	}
	if result.S2A.Monday["1130"] != "CHEM101" || result.S2A.Thursday["1130"] != "CHEM101" {
		t.Error("Error: BaseTimeslotMap did not map course successfully")
	}
	if result.S2A.Tuesday["1130"] != "" || result.S2A.Wednesday["1130"] != "" || result.S2A.Friday["1130"] != "" {
		t.Error("Error: Course added to wrong days")
	}
	if result.S2B.Monday["1130"] != "" {
		t.Error("Error: Course added to wrong stream map")
	}
}

func TestRandomTimeslotAssignment(t *testing.T) {

	testCourse := structs.Course{
		CourseNumber:   "101",
		Subject:        "CHEM",
		SequenceNumber: "A01",
		StreamSequence: "2A",
		CourseTitle:    "Properties of Materials",
	}

	testSchedule := structs.Schedule{
		FallCourses:   []structs.Course{testCourse},
		SpringCourses: []structs.Course{},
		SummerCourses: []structs.Course{},
	}

	testStreamtype := scheduling.CreateEmptyStreamType()
	var err error

	testSchedule.FallCourses, testStreamtype, err = scheduling.AddCoursesToStreamMaps(testSchedule.FallCourses, testStreamtype)
	isAdded := false

	if err != nil {
		t.Error(err)
	}
	if testCourse.Assignment.BeginTime != "" && testCourse.Assignment.EndTime != "" {
		t.Error("Error: Course was not assigned to a block")
	}

	for _, course := range testStreamtype.S2A.Monday {
		if course == "CHEM101" {
			isAdded = true
		}
	}
	for _, course := range testStreamtype.S2A.Tuesday {
		if course == "CHEM101" {
			isAdded = true
		}
	}

	if !isAdded {
		t.Error("Error: Course not added to timeslot map")
	}
}

func TestCantAddConflictingRequiredCourse(t *testing.T) {
	// Test data
	testAssignment := structs.Assignment{
		StartDate: "Sep 05, 2018",
		EndDate:   "Dec 05, 2018",
		BeginTime: "1300",
		EndTime:   "1420",
		HoursWeek: 3,
		Sunday:    false,
		Monday:    true,
		Tuesday:   false,
		Wednesday: false,
		Thursday:  true,
		Friday:    false,
		Saturday:  false,
	}

	testCourse := structs.Course{
		Assignment:     testAssignment,
		CourseNumber:   "101",
		Subject:        "CHEM",
		SequenceNumber: "A01",
		StreamSequence: "2A",
		CourseTitle:    "Properties of Materials",
	}

	testSchedule := structs.Schedule{
		FallCourses:   []structs.Course{testCourse, testCourse},
		SpringCourses: []structs.Course{},
		SummerCourses: []structs.Course{},
	}

	_, err := scheduling.BaseTimeslotMaps(testSchedule.FallCourses)

	if err == nil {
		t.Error("Error: Did not catch required course conflict error")
	}
}

func TestCantScheduleClassOutsideTime(t *testing.T) {
	// Test data
	testAssignment := structs.Assignment{
		StartDate: "Sep 05, 2018",
		EndDate:   "Dec 05, 2018",
		BeginTime: "0300",
		EndTime:   "0420",
		HoursWeek: 3,
		Sunday:    false,
		Monday:    true,
		Tuesday:   false,
		Wednesday: false,
		Thursday:  true,
		Friday:    false,
		Saturday:  false,
	}

	testCourse := structs.Course{
		Assignment:     testAssignment,
		CourseNumber:   "101",
		Subject:        "CHEM",
		SequenceNumber: "A01",
		StreamSequence: "2A",
		CourseTitle:    "Properties of Materials",
	}

	testSchedule := structs.Schedule{
		FallCourses:   []structs.Course{testCourse},
		SpringCourses: []structs.Course{},
		SummerCourses: []structs.Course{},
	}

	_, err := scheduling.BaseTimeslotMaps(testSchedule.FallCourses)

	if err == nil {
		t.Error("Error: Did not catch course being scheduled in improper slot")
	}
}

func TestFullRandomAssignment(t *testing.T) {
	allChanged := true

	jsonFile, err := os.Open("../data/base-courses-test.json")
	if err != nil {
		t.Error("Error: Test file not found")
	}

	courseData, _ := ioutil.ReadAll(jsonFile)

	testSchedule, err := structs.ParseHistorical(courseData)

	if err != nil {
		t.Error("Error: Course data parsing failed")
	}

	testStreamtype := scheduling.CreateEmptyStreamType()

	testSchedule.FallCourses, _, err = scheduling.AddCoursesToStreamMaps(testSchedule.FallCourses, testStreamtype)

	if err != nil {
		t.Error(err)
	}

	for _, course := range testSchedule.FallCourses {
		if course.Assignment.BeginTime == "" || course.Assignment.EndTime == "" {
			allChanged = false
		}
	}

	if !allChanged {
		t.Error("Error: Some courses in list were not assigned times")
	}
}

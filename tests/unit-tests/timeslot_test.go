package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"
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

	result, err := scheduling.BaseTimeslotMaps(testSchedule.FallCourses, "Fall")

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
		NumSections:    1,
		CourseCapacity: 100,
	}

	testSchedule := structs.Schedule{
		FallCourses:   []structs.Course{testCourse},
		SpringCourses: []structs.Course{},
		SummerCourses: []structs.Course{},
	}

	testStreamtype := scheduling.CreateEmptyStreamType()
	var err error

	testSchedule.FallCourses, testStreamtype, err = scheduling.AddCoursesToStreamMaps(scheduling.Split(testSchedule.FallCourses), testStreamtype, "Fall")
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

	_, err := scheduling.BaseTimeslotMaps(testSchedule.FallCourses, "Fall")

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

	_, err := scheduling.BaseTimeslotMaps(testSchedule.FallCourses, "Fall")

	if err == nil {
		t.Error("Error: Did not catch course being scheduled in improper slot")
	}
}

func TestCantScheduleCourseWithoutStreamSequence(t *testing.T) {
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
		StreamSequence: "",
		CourseTitle:    "Properties of Materials",
	}

	testSchedule := structs.Schedule{
		FallCourses:   []structs.Course{testCourse},
		SpringCourses: []structs.Course{},
		SummerCourses: []structs.Course{},
	}

	_, err := scheduling.BaseTimeslotMaps(testSchedule.FallCourses, "Fall")

	if err == nil {
		t.Error("Error: Did not catch course being scheduled without stream sequence")
	}
}

func TestStreamOverflow(t *testing.T) {
	// Test data
	jsonFile, err := os.Open("../data/overflow-courses-test.json")
	if err != nil {
		t.Error("Error: Test file not found")
	}

	courseData, _ := ioutil.ReadAll(jsonFile)

	testSchedule, err := structs.ParseHistorical(courseData)
	if err != nil {
		t.Error("Error: Parsing data from JSON to schedule object failed")
	}

	_, err = scheduling.BaseTimeslotMaps(testSchedule.FallCourses, "Fall")

	if err == nil {
		t.Error("Error: Did not catch course being scheduled without stream sequence")
	}
}

func TestSetDate(t *testing.T) {
	// Test data
	testAssignment := structs.Assignment{
		StartDate: "",
		EndDate:   "",
		BeginTime: "",
		EndTime:   "",
		HoursWeek: 0,
		Sunday:    false,
		Monday:    false,
		Tuesday:   false,
		Wednesday: false,
		Thursday:  false,
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
		SpringCourses: []structs.Course{testCourse},
		SummerCourses: []structs.Course{testCourse},
	}

	testFallStreamType := scheduling.CreateEmptyStreamType()
	testSpringStreamType := scheduling.CreateEmptyStreamType()
	testSummerStreamType := scheduling.CreateEmptyStreamType()

	fallCourses, _, err := scheduling.AddCoursesToStreamMaps(testSchedule.FallCourses, testFallStreamType, "Fall")
	if err != nil {
		t.Error(err)
	}

	springCourses, _, err := scheduling.AddCoursesToStreamMaps(testSchedule.SpringCourses, testSpringStreamType, "Spring")
	if err != nil {
		t.Error(err)
	}

	summerCourses, _, err := scheduling.AddCoursesToStreamMaps(testSchedule.SummerCourses, testSummerStreamType, "Summer")
	if err != nil {
		t.Error(err)
	}

	year := time.Now().Year()

	if fallCourses[0].Assignment.StartDate != "Sep 01, "+strconv.Itoa(year) {
		t.Errorf("error: incorrect date assigned, expected Sep 01, %v, got %v", year, fallCourses[0].Assignment.StartDate)
	}
	if fallCourses[0].Assignment.EndDate != "Dec 01, "+strconv.Itoa(year) {
		t.Errorf("error: incorrect date assigned, expected Dec 01, %v, got %v", year, fallCourses[0].Assignment.EndDate)
	}

	if springCourses[0].Assignment.StartDate != "Jan 01, "+strconv.Itoa(year+1) {
		t.Errorf("error: incorrect date assigned, expected Jan 01, %v, got %v", year+1, springCourses[0].Assignment.StartDate)
	}
	if springCourses[0].Assignment.EndDate != "Apr 01, "+strconv.Itoa(year+1) {
		t.Errorf("error: incorrect date assigned, expected Apr 01, %v, got %v", year+1, springCourses[0].Assignment.EndDate)
	}

	if summerCourses[0].Assignment.StartDate != "May 01, "+strconv.Itoa(year+1) {
		t.Errorf("error: incorrect date assigned, expected May 01, %v, got %v", year+1, summerCourses[0].Assignment.StartDate)
	}
	if summerCourses[0].Assignment.EndDate != "Aug 01, "+strconv.Itoa(year+1) {
		t.Errorf("error: incorrect date assigned, expected Aug 01, %v, got %v", year+1, summerCourses[0].Assignment.EndDate)
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

	testSchedule.FallCourses, _, err = scheduling.AddCoursesToStreamMaps(scheduling.Split(testSchedule.FallCourses), testStreamtype, "Fall")

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

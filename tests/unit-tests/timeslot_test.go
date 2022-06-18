package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
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

	if err != "" {
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
		FallCourses:   []structs.Course{testCourse, testCourse},
		SpringCourses: []structs.Course{},
		SummerCourses: []structs.Course{},
	}

	_, err := scheduling.BaseTimeslotMaps(testSchedule.FallCourses)

	if err == "" {
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

	_, err := scheduling.BaseTimeslotMaps(testSchedule.FallCourses)

	if err == "" {
		t.Error("Error: Did not catch course being scheduled in improper slot")
	}
}
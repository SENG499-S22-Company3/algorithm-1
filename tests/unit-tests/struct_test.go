package tests

import (
	"algorithm-1/structs"
	"testing"
)

// Test course info
var courseNum uint = 370
var courseSub = "CSC"
var courseSeq = "A01/A02"
var courseTitle = "Database Systems"
var courseEquipment = []string{"Projector", "Whiteboard"}

// Test professor info
var displayName = "John Doe"
var profEquipment = []string{"Projector"}

func TestCourseStruct(t *testing.T) {
	testCourse := structs.Course{
		CourseNumber:      courseNum,
		Subject:           courseSub,
		SequenceNumber:    courseSeq,
		CourseTitle:       courseTitle,
		RequiredEquipment: courseEquipment,
	}

	if testCourse.CourseNumber != 370 {
		t.Errorf("Got %v. expected 370", testCourse.CourseNumber)
	}
}

func TestProfessorStuct(t *testing.T) {
	testPreference := structs.Preference{
		CourseNum:     "CSC370",
		PreferenceNum: 195,
	}

	testProfessor := structs.Professor{
		Preferences:       []structs.Preference{testPreference},
		DisplayName:       displayName,
		RequiredEquipment: profEquipment,
	}

	if testProfessor.Preferences[0].CourseNum != "CSC370" {
		t.Errorf("Got %v, expected %v", testProfessor.Preferences[0].CourseNum, 370)
	}
}

func TestAssignmentStruct(t *testing.T) {
	// testCourse := structs.Course{}
	// testAssignment := structs.Assignment{}

}

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
var requiresPEng = false

// Test professor info
var displayName = "John Doe"
var numCourses uint = 3
var profEquipment = []string{"Projector"}
var hasPEng = true

func TestCourseStruct(t *testing.T) {
	testCourse := structs.Course{
		CourseNumber:      courseNum,
		Subject:           courseSub,
		SequenceNumber:    courseSeq,
		CourseTitle:       courseTitle,
		RequiredEquipment: courseEquipment,
		RequiresPEng:      requiresPEng,
	}

	if testCourse.CourseNumber != 370 {
		t.Errorf("Got %v. expected 370", testCourse.CourseNumber)
	}
}

func TestProfessorStuct(t *testing.T) {
	testCourse := structs.Course{
		CourseNumber:      courseNum,
		Subject:           courseSub,
		SequenceNumber:    courseSeq,
		CourseTitle:       courseTitle,
		RequiredEquipment: courseEquipment,
		RequiresPEng:      requiresPEng,
	}

	testPreference := structs.Preference{
		Course:        testCourse,
		PreferenceNum: 195,
	}

	testProfessor := structs.Professor{
		Preferences:        []structs.Preference{testPreference},
		CoursesCanTeach:    []structs.Course{testCourse},
		DisplayName:        displayName,
		NumCoursesCanTeach: numCourses,
		RequiredEquipment:  profEquipment,
		HasPEng:            hasPEng,
	}

	if testProfessor.Preferences[0].Course.CourseNumber != 370 {
		t.Errorf("Got %v, expected %v", testProfessor.Preferences[0].Course.CourseNumber, 370)
	}
}

func TestAssignmentStruct(t *testing.T) {
	// testCourse := structs.Course{}
	// testAssignment := structs.Assignment{}

}

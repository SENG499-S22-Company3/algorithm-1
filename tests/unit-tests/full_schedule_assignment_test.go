package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func getInput(t *testing.T) (structs.Schedule, structs.Input){
	jsonData, err := ioutil.ReadFile("../data/input-test.json")
	if err != nil {
		log.Fatal("Error when opening input-test.json file: ", err)
	}

	input, err := structs.ParseInput(jsonData)
	if err != nil {
		t.Error("Input parsing failed with error: ", err.Error())
	}

	if input.HardScheduled.SpringCourses == nil {
		t.Error("Input failed to be parsed: fall historical courses should not be null")
	}

	jsonFile, err := os.Open("../data/base-courses-test.json")
	if err != nil {
		t.Error("Error: Test file not found")
	}

	courseData, _ := ioutil.ReadAll(jsonFile)
	testSchedule, err := structs.ParseHistorical(courseData)
	if err != nil {
		t.Error("Error: Course data parsing failed")
	}
	return testSchedule, input
}

func printAssignments(testScheduleCourse []structs.Course, prefsMap map[string]map[string]int){
	for _,c := range testScheduleCourse{
		fmt.Println(c.CourseTitle, c.SequenceNumber,"in sequence", c.StreamSequence)
		fmt.Println("\t taught by:", c.Prof.DisplayName, "( preference:", prefsMap[c.Prof.DisplayName][c.Subject+c.CourseNumber],")" )
		fmt.Println("\t\t at", c.Assignment.BeginTime ,"to",c.Assignment.EndTime )
		if(c.Assignment.Monday == true){
			fmt.Println("\t\t\t on MTh")
		}else {
			fmt.Println("\t\t\t on TWF")
		}
	}
}

func TestFallScheduleAssignment(t *testing.T) {

	testSchedule, input := getInput(t)
	testStreamtype := scheduling.CreateEmptyStreamType()
	testSchedule.FallCourses, _, _ = scheduling.AddCoursesToStreamMaps(testSchedule.FallCourses, testStreamtype)
	testScheduleCourse := scheduling.AssignCourseProf(input.HardScheduled.FallCourses, testSchedule.FallCourses, input.Professors, "Fall")
	err := scheduling.ScheduleConstraintsCheck("Fall", testScheduleCourse, input)

	if err != nil {
		t.Error(err)
	}
}

func TestSpringScheduleAssignment(t *testing.T) {

	testSchedule, input := getInput(t)
	testStreamtype := scheduling.CreateEmptyStreamType()
	testSchedule.SpringCourses, _, _ = scheduling.AddCoursesToStreamMaps(testSchedule.SpringCourses, testStreamtype)
	testScheduleCourse := scheduling.AssignCourseProf(input.HardScheduled.SpringCourses, testSchedule.SpringCourses, input.Professors, "Spring")
	err := scheduling.ScheduleConstraintsCheck("Spring", testScheduleCourse, input)

	if err != nil {
		t.Error(err)
	}
}

func TestSummerScheduleAssignment(t *testing.T) {

	testSchedule, input := getInput(t)
	testStreamtype := scheduling.CreateEmptyStreamType()
	testSchedule.SummerCourses, _, _ = scheduling.AddCoursesToStreamMaps(testSchedule.SummerCourses, testStreamtype)
	testScheduleCourse := scheduling.AssignCourseProf(input.HardScheduled.SummerCourses, testSchedule.SummerCourses, input.Professors, "Summer")
	err := scheduling.ScheduleConstraintsCheck("Summer", testScheduleCourse, input)

	if err != nil {
		t.Error(err)
	}
}
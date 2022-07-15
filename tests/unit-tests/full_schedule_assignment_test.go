package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func getInput(t *testing.T) (structs.Schedule, structs.Input) {
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

func TestFallScheduleAssignment(t *testing.T) {

	testSchedule, input := getInput(t)
	testStreamtype, err := scheduling.BaseTimeslotMaps(input.HardScheduled.FallCourses, "Fall")
	if err != nil {
		t.Error(err)
	}
	testSchedule.FallCourses, _, err = scheduling.AddCoursesToStreamMaps(scheduling.Split(testSchedule.FallCourses), testStreamtype, "Fall")
	if err != nil {
		t.Error(err)
	}
	testScheduleCourse := scheduling.AssignCourseProf(input.HardScheduled.FallCourses, testSchedule.FallCourses, input.Professors, "Fall")
	testScheduleCourse = append(testScheduleCourse, input.HardScheduled.FallCourses...)

	err = scheduling.ScheduleConstraintsCheck("Fall", testScheduleCourse, input.Professors)
	if err != nil {
		t.Error(err)
	}

	_, err = scheduling.BaseTimeslotMaps(testScheduleCourse, "Fall")
	if err != nil {
		t.Error(err)
	}
}

func TestSpringScheduleAssignment(t *testing.T) {

	testSchedule, input := getInput(t)
	testStreamtype, err := scheduling.BaseTimeslotMaps(input.HardScheduled.SpringCourses, "Spring")
	if err != nil {
		t.Error(err)
	}
	testSchedule.SpringCourses, _, err = scheduling.AddCoursesToStreamMaps(scheduling.Split(testSchedule.SpringCourses), testStreamtype, "Spring")
	if err != nil {
		t.Error(err)
	}
	testScheduleCourse := scheduling.AssignCourseProf(input.HardScheduled.SpringCourses, testSchedule.SpringCourses, input.Professors, "Spring")
	testScheduleCourse = append(testScheduleCourse, input.HardScheduled.SpringCourses...)

	err = scheduling.ScheduleConstraintsCheck("Spring", testScheduleCourse, input.Professors)
	if err != nil {
		t.Error(err)
	}

	_, err = scheduling.BaseTimeslotMaps(testScheduleCourse, "Spring")
	if err != nil {
		t.Error(err)
	}
}

func TestSummerScheduleAssignment(t *testing.T) {

	testSchedule, input := getInput(t)
	testStreamtype, err := scheduling.BaseTimeslotMaps(input.HardScheduled.SummerCourses, "Summer")
	if err != nil {
		t.Error(err)
	}
	testSchedule.SummerCourses, _, err = scheduling.AddCoursesToStreamMaps(scheduling.Split(testSchedule.SummerCourses), testStreamtype, "Summer")
	if err != nil {
		t.Error(err)
	}
	testScheduleCourse := scheduling.AssignCourseProf(input.HardScheduled.SummerCourses, testSchedule.SummerCourses, input.Professors, "Summer")
	testScheduleCourse = append(testScheduleCourse, input.HardScheduled.SummerCourses...)

	err = scheduling.ScheduleConstraintsCheck("Summer", testScheduleCourse, input.Professors)
	if err != nil {
		t.Error(err)
	}

	_, err = scheduling.BaseTimeslotMaps(testScheduleCourse, "Summer")
	if err != nil {
		t.Error(err)
	}
}

func TestTBDScheduleAssignment(t *testing.T) {

	testSchedule, input := getInput(t)
	testStreamtype, err := scheduling.BaseTimeslotMaps(input.HardScheduled.SummerCourses, "Summer")

	testSchedule.SummerCourses = append(testSchedule.SummerCourses, structs.Course{
		CourseNumber:   "225",
		Subject:        "TEST",
		SequenceNumber: "A01",
		CourseTitle:    "Fake Course",
		StreamSequence: "2B",
		NumSections:    2,
		CourseCapacity: 100,
	})

	testSchedule.SummerCourses, _, err = scheduling.AddCoursesToStreamMaps(scheduling.Split(testSchedule.SummerCourses), testStreamtype, "Summer")
	if err != nil {
		t.Error(err)
	}
	testScheduleCourse := scheduling.AssignCourseProf(input.HardScheduled.SummerCourses, testSchedule.SummerCourses, input.Professors, "Summer")
	testScheduleCourse = append(testScheduleCourse, input.HardScheduled.SummerCourses...)

	err = scheduling.ScheduleConstraintsCheck("Summer", testScheduleCourse, input.Professors)
	if err != nil {
		t.Error(err)
	}
	_, err = scheduling.BaseTimeslotMaps(testScheduleCourse, "Summer")
	if err != nil {
		t.Error(err)
	}
}

func TestTeachingMax(t *testing.T) {
	jsonData, err := ioutil.ReadFile("../data/teaching-max-test.json")
	if err != nil {
		log.Fatal("Error when opening input-test.json file: ", err)
	}

	input, err := structs.ParseInput(jsonData)
	if err != nil {
		t.Error("Input parsing failed with error: ", err.Error())
	}

	courses := scheduling.Assignments(input.HardScheduled.FallCourses, input.CoursesToSchedule.FallCourses, input.Professors, "Fall")
	if err != nil {
		t.Error(err)
	}

	if len(courses) != 2 {
		t.Error("Courses slice should be of len 2 not size ", len(courses))
	}

	count := 0
	for _, c := range courses {
		if c.Prof.DisplayName == "TBD" {
			count++
		}
	}

	if count != 1 {
		t.Error("Bill Bird Teaching too many courses")
	}

}

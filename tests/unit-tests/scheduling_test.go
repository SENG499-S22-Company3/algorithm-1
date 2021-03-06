package tests

import (
	"algorithm-1/genetic"
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// finds a course in a list given string like "CSC 115"
func CheckCourse(id string, list []structs.Course) bool {
	tokens := strings.Split(id, " ")
	for _, b := range list {
		if b.Subject == tokens[0] && b.CourseNumber == tokens[1] {
			return true
		}
	}
	return false
}

func TestSmallScaleBase(t *testing.T) {
	// Test data including 1 course to ignore (chem 101), 1 course to include (csc 115), and 1 course to exclude (csc 230)
	historicalData := []byte(`
		{
			"fallCourses": [
				{
					"courseNumber": "101",
					"subject": "CHEM",
					"sequenceNumber": "A01",
					"courseTitle": "Properties of Materials",
					"assignment": {
						"beginTime": "1300",
						"endDate": "Dec 05, 2018",
						"endTime": "1420",
						"friday": false,
						"hoursWeek": 3,
						"monday": true,
						"saturday": false,
						"startDate": "Sep 05, 2018",
						"sunday": false,
						"thursday": true,
						"tuesday": false,
						"wednesday": false
					}
				},
				{
					"courseNumber": "115",
					"subject": "CSC",
					"sequenceNumber": "A01",
					"courseTitle": "FUNDAMENTAL PROGRAMING:II",
					"assignment": {
						"startDate": "Sep 03, 2014",
						"endDate": "Dec 03, 2014",
						"beginTime": "1530",
						"endTime": "1620",
						"hoursWeek": 3,
						"sunday": false,
						"monday": true,
						"tuesday": false,
						"wednesday": true,
						"thursday": true,
						"friday": false,
						"saturday": false
					}
				},
				{
					"courseNumber": "230",
					"subject": "CSC",
					"sequenceNumber": "A02",
					"courseTitle": "COMPUTER ARCHITECTURE",
					"assignment": {
						"startDate": "Sep 03, 2014",
						"endDate": "Dec 03, 2014",
						"beginTime": "1300",
						"endTime": "1420",
						"hoursWeek": 3,
						"sunday": false,
						"monday": true,
						"tuesday": false,
						"wednesday": false,
						"thursday": true,
						"friday": false,
						"saturday": false
					}
				  }
			],
			"springCourses": [],
			"summerCourses": []
		}`)

	historicalResult, _ := structs.ParseHistorical(historicalData)
	historical := historicalResult.FallCourses

	// the course that we want to include
	courseData := []byte(
		`[{
			"streamSequence": "3A",
			"courseNumber": "115",
			"subject": "CSC",
			"sequenceNumber": "A01",
			"courseTitle": "FUNDAMENTAL PROGRAMING:II"
		}]`)
	testCourses, _ := structs.ParseCourses(courseData)

	result := scheduling.BaseSemester(testCourses, historical)

	if len(result) != 2 {
		t.Error("Incorrect number of courses")
	}
	if result[0].CourseNumber != "101" {
		t.Error("Chem 101 should be included")
	}
	if result[1].CourseNumber != "115" {
		t.Error("The requested course should be included")
	}
}

func TestBaseSchedule(t *testing.T) {
	jsonFile, err := os.Open("../data/historical-data-2019-test.json")
	if err != nil {
		t.Error("File not found")
	}
	historicalData, _ := ioutil.ReadAll(jsonFile) // making byte array

	historicalResult, _ := structs.ParseHistorical(historicalData)
	historical := historicalResult.FallCourses
	jsonFile.Close()

	jsonFile, err = os.Open("../data/base-courses-test.json")
	if err != nil {
		t.Error("File not found")
	}
	courseData, _ := ioutil.ReadAll(jsonFile) // making byte array
	courseResult, _ := structs.ParseHistorical(courseData)
	testCourses := courseResult.FallCourses
	jsonFile.Close()

	result := scheduling.BaseSemester(testCourses, historical)
	if !CheckCourse("SENG 265", result) {
		t.Error("Missing course")
	}
	if len(result) != 16 {
		t.Error("Incorrect number of courses")
	}
}

func TestBaseScheduleConcurrent(t *testing.T) {
	jsonFile, err := os.Open("../data/historical-data-2019-test.json")
	if err != nil {
		t.Error("File not found")
	}
	historicalData, _ := ioutil.ReadAll(jsonFile) // making byte array
	historical, _ := structs.ParseHistorical(historicalData)
	jsonFile.Close()
	jsonFile, err = os.Open("../data/base-courses-test.json")
	if err != nil {
		t.Error("File not found")
	}
	courseData, _ := ioutil.ReadAll(jsonFile) // making byte array
	testCourses, _ := structs.ParseHistorical(courseData)
	jsonFile.Close()

	schedule := scheduling.BaseSchedule(testCourses, historical)

	// fmt.Print(len(schedule.FallCourses))
	// fmt.Print(len(schedule.SummerCourses))
	// fmt.Print(len(schedule.SpringCourses))

	// verify
	if len(schedule.FallCourses) != 16 {
		t.Error("Incorrect number of fall courses")
	}
	if len(schedule.SpringCourses) != 52 {
		t.Error("Incorrect number of spring courses")
	}
	if len(schedule.SummerCourses) != 24 {
		t.Error("Incorrect number of summer courses")
	}
	if !CheckCourse("ECON 180", schedule.SummerCourses) {
		t.Error("Missing course")
	}
	if !CheckCourse("CSC 115", schedule.SummerCourses) {
		t.Error("Missing course")
	}
	if CheckCourse("SENG 499", schedule.SummerCourses) {
		t.Error("Course should not be present")
		// Shouldn't be included because seng courses aren't present in historical data for summer 2019
	}
}

func TestGeneticAlg(t *testing.T) {
	var input structs.Input
	term := "Spring"

	jsonData, err := ioutil.ReadFile("../data/input-test.json")
	if err != nil {
		t.Error("Error when opening input-test.json file: ", err)
	}

	input, err = structs.ParseInput(jsonData)
	if err != nil {
		t.Error("Input parsing failed with error: ", err.Error())
	}

	result, _ := genetic.RunGeneticAlg(input.HardScheduled.SpringCourses, input.CoursesToSchedule.SpringCourses, input.Professors, term)
	if err != nil {
		t.Error(err)
	}

	for _, course := range result {
		fmt.Printf("%+v%+v  %+v  %+v %+v %+v\n", course.Subject, course.CourseNumber, course.StreamSequence, course.Assignment.BeginTime, course.Assignment.EndTime, course.Prof.DisplayName)
	}

}

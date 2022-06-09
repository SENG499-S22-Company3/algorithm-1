package tests

import (
	"algorithm-1/structs"
	"strings"
	"testing"
)

func TestSmallHistoricalParse(t *testing.T) {
	jsonData := []byte(`
		{
			"fallTermCourses": [
				{
				"courseNumber": "101",
				"subject": "CHEM",
				"sequenceNumber": "A01",
				"courseTitle": "Properties of Materials",
				"meetingTime": {
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
					"courseNumber": "101",
					"subject": "CHEM",
					"sequenceNumber": "A01",
					"courseTitle": "Properties of Materials",
					"meetingTime": {
						"beginTime": "1300",
						"endDate": "Dec 05, 2017",
						"endTime": "1420",
						"friday": false,
						"hoursWeek": 3,
						"monday": true,
						"saturday": false,
						"startDate": "Sep 05, 2017",
						"sunday": false,
						"thursday": true,
						"tuesday": false,
						"wednesday": false
					}
					}
			],
			"springTermCourses": [],
			"summerTermCourses": []
		}`)

	testSchedule := structs.ParseHistorical(jsonData)

	if testSchedule.FallCourses == nil {
		t.Error("Schedule failed to be parsed")
	} else if testSchedule.FallCourses[0].CourseNumber != "101" && !testSchedule.FallCourses[0].Assignment.Thursday {
		t.Errorf("Schedule successfully parsed, but data is incorrect. Course number should be 101 and it was %v, and/or Thursday should be true when it was %v",
			testSchedule.FallCourses[0].CourseNumber, testSchedule.FallCourses[0].Assignment.Thursday)
	}

	schedule2019 := structs.Schedule2019(testSchedule)

	if schedule2019.FallCourses == nil {
		t.Error("Schedule failed to be parsed")
	} else if len(schedule2019.FallCourses) != 1 {
		t.Errorf("Schedule successfully parsed, but len not")
	}
}

func TestLargeHistoricalParse(t *testing.T) {
	var testSchedule structs.Schedule

	jsonData := []byte(`
		{
			"fallTermCourses": [],
			"springTermCourses": [],
			"summerTermCourses": []
		}`)

	testSchedule = structs.ParseHistorical(jsonData)

	if testSchedule.FallCourses == nil {
		t.Error("Schedule failed to be parsed")
	}
}

func TestCourseParse(t *testing.T) {
	// example input
	jsonData := []byte(
		`[{
			"streamSequence": "3A",
			"courseNumber": "115",
			"subject": "CSC",
			"sequenceNumber": "A01",
			"courseTitle": "FUNDAMENTAL PROGRAMING:II"
		}, {
			"streamSequence": "3B",
			"courseNumber": "225",
			"subject": "CSC",
			"sequenceNumber": "A01",
			"courseTitle": "ALGORITHMS+DATA STUCT:I"
		}]`)
	result := structs.ParseCourses(jsonData)

	if result[0].CourseNumber != "115" {
		t.Error("Incorrect CourseNumber")
	}
	if result[1].CourseNumber != "225" {
		t.Error("Incorrect CourseNumber")
	}
}

func TestJSONGeneration(t *testing.T) {
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
		FallCourses:   []structs.Course{testCourse},
		SpringCourses: []structs.Course{},
		SummerCourses: []structs.Course{},
	}

	// Test check data with whitespace for readability
	jsonString := `
	{
		"fallTermCourses":
		[
			{
				"courseNumber":"101",
				"subject":"CHEM",
				"sequenceNumber":"A01",
				"courseTitle":"Properties of Materials",
				"streamSequence":"2A",
				"meetingTime":
					{
					"startDate":"Sep 05, 2018",
					"endDate":"Dec 05, 2018",
					"beginTime":"1300",
					"endtime":"1420",
					"hoursWeek":3,
					"sunday":false,
					"monday":true,
					"tuesday":false,
					"wednesday":false,
					"thursday":true,
					"friday":false,
					"saturday":false
					},
				"prof":
					{
					"displayName":"JohnSmith"
					}
			}
		],
		"springTermCourses":[],
		"summerTermCourses":[]
	}`

	jsonData := structs.StructToJSON(testSchedule)

	// Remove whitespace and newlines for testing
	jsonString = strings.Replace(jsonString, "\n", "", -1)
	jsonString = strings.Replace(jsonString, "\t", "", -1)

	if string(jsonData) != jsonString {
		t.Error("Schedule failed to parse to JSON correctly")
	}
}

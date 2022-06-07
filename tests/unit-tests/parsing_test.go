package tests

import (
	"algorithm-1/structs"
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
					"beginTime": 1300,
					"endDate": "Dec 05, 2018",
					"endTime": 1420,
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

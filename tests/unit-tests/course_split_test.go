package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"testing"
)

func TestSplitCourses(t *testing.T) {
	// Setup
	jsonData := []byte(
		`[{
      "courseNumber": "111",
      "subject": "CSC",
      "sequenceNumber": "A01",
      "streamSequence": "1A",
      "courseTitle": "Fundamentals of Programming with Engineering Applications",
      "assignment": {
        "startDate": "Jan 07, 2019",
        "endDate": "Apr 05, 2019",
        "beginTime": "0830",
        "endTime": "0930",
        "hoursWeek": 3,
        "sunday": false,
        "monday": true,
        "tuesday": false,
        "wednesday": false,
        "thursday": true,
        "friday": false,
        "saturday": false
      },
      "prof": {
        "preferences": [
          {
            "courseNum": "CSC111",
            "preferenceNum": 0,
            "term": "FALL"
          }
        ],
        "displayName": "Michael, Zastre",
        "fallTermCourses": 1,
        "springTermCourses": 1,
        "summerTermCourses": 1
      },
      "courseCapacity": 400,
      "numSections": 0
    },
    {
      "courseNumber": "116",
      "subject": "CSC",
      "sequenceNumber": "A01",
      "streamSequence": "1A",
      "courseTitle": "Fundamentals of Programming with Engineering Applications",
      "assignment": {
        "startDate": "Jan 07, 2019",
        "endDate": "Apr 05, 2019",
        "beginTime": "0830",
        "endTime": "0930",
        "hoursWeek": 3,
        "sunday": false,
        "monday": true,
        "tuesday": false,
        "wednesday": false,
        "thursday": true,
        "friday": false,
        "saturday": false
      },
      "prof": {
        "preferences": [
          {
            "courseNum": "CSC111",
            "preferenceNum": 0,
            "term": "FALL"
          }
        ],
        "displayName": "Michael, Zastre",
        "fallTermCourses": 1,
        "springTermCourses": 1,
        "summerTermCourses": 1
      },
      "courseCapacity": 600,
      "numSections": 0
    }]`)

	courses, err := structs.ParseCourses(jsonData)
	if err != nil {
		t.Error("Parsing courses failed with error: ", err.Error())
	}

	courses = scheduling.Split(courses)

	if len(courses) != 10 {
		t.Error("Courses slice should be of len 10 not size ", len(courses))
	}

}

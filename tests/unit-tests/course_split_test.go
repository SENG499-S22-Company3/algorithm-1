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
          "courseNumber": "265",
          "subject": "SENG",
          "sequenceNumber": "A01",
          "courseTitle": "Intro to Software Development",
          "StreamSequence": "2B",
          "numSections": 0,
          "courseCapacity": 477
        },
        {
          "courseNumber": "226",
          "subject": "SENG",
          "sequenceNumber": "A01",
          "courseTitle": "Algorithms and Data Structures II",
          "StreamSequence": "3A",
          "numSections": 1,
          "courseCapacity": 225
          },
        {
          "courseNumber": "225",
          "subject": "CSC",
          "sequenceNumber": "A01",
          "courseTitle": "Algorithms and Data Structures I",
          "StreamSequence": "2B",
          "numSections": 2,
          "courseCapacity": 345
        },
        {
          "courseNumber": "111",
          "subject": "CSC",
          "sequenceNumber": "A01",
          "courseTitle": "Fundamentals of Programming with Engineering Applications",
          "StreamSequence": "1A",
          "numSections": 3,
          "courseCapacity": 436
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

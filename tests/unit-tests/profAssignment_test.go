package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestProfAssignment(t *testing.T) {
	jsonData, err := ioutil.ReadFile("../data/input-test.json")
    if err != nil {
        log.Fatal("Error when opening input-test.json file: ", err)
    }
	
	input, err := structs.ParseInput(jsonData)
	if err != nil {
		t.Error("Input parsing failed with error: ", err.Error())
	}

	if input.HistoricData.FallCourses == nil {
		t.Error("Input failed to be parsed: fall historical courses should not be null")
	}

	testScheduleCourse := scheduling.AssignCourseProf(input.HistoricData.FallCourses, input.CoursesToSchedule.FallCourses, input.Professors)

	for _,c := range testScheduleCourse {
		if c.Prof.DisplayName == "" {
			t.Error("Professors not assigned to course")
		}
	}

	for _,c := range testScheduleCourse{
		fmt.Println(c.CourseTitle, "in sequence", c.StreamSequence)
		fmt.Println("\t taught by:", c.Prof.DisplayName)
		fmt.Println("\t\t at", c.Assignment.BeginTime ,"to",c.Assignment.EndTime )
		if(c.Assignment.Monday == true){
			fmt.Println("\t\t\t on MTh")
		}else {
			fmt.Println("\t\t\t on TWF")
		}
	}
	
}
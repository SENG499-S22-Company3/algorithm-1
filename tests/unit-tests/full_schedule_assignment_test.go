package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestFullScheduleAssignment(t *testing.T) {
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

	jsonFile, err := os.Open("../data/base-courses-test.json")
	if err != nil {
		t.Error("Error: Test file not found")
	}

	courseData, _ := ioutil.ReadAll(jsonFile)
	testSchedule, err := structs.ParseHistorical(courseData)
	if err != nil {
		t.Error("Error: Course data parsing failed")
	}

	testStreamtype := scheduling.CreateEmptyStreamType()
	testSchedule.FallCourses, _, err = scheduling.AddCoursesToStreamMaps(testSchedule.FallCourses, testStreamtype)
	testScheduleCourse := scheduling.AssignCourseProf(input.HistoricData.FallCourses, testSchedule.FallCourses, input.Professors)

	// for _,c := range testSchedule.FallCourses{
	// 	fmt.Println(c.CourseTitle, "in sequence", c.StreamSequence)
	// 	fmt.Println("\t taught by:", c.Prof.DisplayName)
	// 	fmt.Println("\t\t at", c.Assignment.BeginTime ,"to",c.Assignment.EndTime )
	// 	if(c.Assignment.Monday == true){
	// 		fmt.Println("\t\t\t on MTh")
	// 	}else {
	// 		fmt.Println("\t\t\t on TWF")
	// 	}
	// }

	for _,c := range testScheduleCourse {
		if c.Prof.DisplayName == "" ||  c.Assignment.BeginTime == "" || c.Assignment.EndTime == ""{
			t.Error("Schedule not properlly assigned")
		}
	}
}

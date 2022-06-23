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

	if input.HistoricData.SpringCourses == nil {
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
	testSchedule.SpringCourses, _, err = scheduling.AddCoursesToStreamMaps(testSchedule.SpringCourses, testStreamtype)
	testScheduleCourse := scheduling.AssignCourseProf(input.HistoricData.SpringCourses, testSchedule.SpringCourses, input.Professors)
	
	var teachingMap = map[string]map[string]string{}

	for _,p := range input.Professors{
		teachingMap[p.DisplayName] = map[string]string{}
	}

	// profsMap, _ := scheduling.MapPreferences(input.Professors)
	// fmt.Println("# of courses:", len(testScheduleCourse))
	// for i,c := range testScheduleCourse{
	// 	fmt.Println(i, c.CourseTitle, "in sequence", c.StreamSequence)
	// 	fmt.Println("\t taught by:", c.Prof.DisplayName, "( preference:" ,profsMap[c.Prof.DisplayName][c.Subject+c.CourseNumber],")" )
	// 	fmt.Println("\t\t at", c.Assignment.BeginTime ,"to",c.Assignment.EndTime )
	// 	if(c.Assignment.Monday == true){
	// 		fmt.Println("\t\t\t on MTh")
	// 	}else {
	// 		fmt.Println("\t\t\t on TWF")
	// 	}
	// }

	for _,c := range testScheduleCourse {
		var d string
		if(c.Assignment.Monday == true){
			d = "MTh"+c.Assignment.BeginTime
		} else {
			d = "TWF"+c.Assignment.BeginTime
		}
		
		if c.Prof.DisplayName == "" ||  c.Assignment.BeginTime == "" || c.Assignment.EndTime == ""{
			t.Error("Error: Schedule not properlly assigned")
		}

		if _, err := teachingMap[c.Prof.DisplayName][d]; err {
			t.Error("Error: Prof teaching another course at this time.")
		}

		teachingMap[c.Prof.DisplayName][d] = c.CourseTitle+d
	}
}

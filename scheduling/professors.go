package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"math/rand"
	"time"
)

func randomizer(profList []string) []string{
	rand.Seed(time.Now().UnixMilli())
	for i := range profList {
		j := rand.Intn(i + 1)
		profList[i], profList[j] = profList[j], profList[i]
	}
	return profList
}

/*
	Input:  profs []structs.Professor
	Output: profsMap map[string]map[string]int, profList []string
*/
func mapPreferences(profs []structs.Professor) (map[string]map[string]int, []string){
	var profsMap = map[string]map[string]int{}
	var profList []string

	for _, s := range profs {
		profsMap[s.DisplayName] = map[string]int{}
		profList = append(profList, s.DisplayName)
		for _, x := range s.Preferences {
			profsMap[s.DisplayName][x.CourseNum] = int(x.PreferenceNum)
		}
	}

	return profsMap, randomizer(profList)
}

/*
	Input:  profsMap map[string]map[string]int, profList []string, teachingMap map[string]map[string]string, course string
	Output: prof string
*/
func assignProf(profsMap map[string]map[string]int, profList []string, teachingMap map[string]map[string]string, course structs.Course) (string){
	var max int = 0
	var prof string = "N/A"

	var t = course.Assignment.BeginTime
	var d = "MTh"+t
	if course.Assignment.Monday == true{
		d = "TWF"+t
	}
	var c = course.Subject+course.CourseNumber

	for _, p := range profList {
		if max < profsMap[p][c]{
			// make sure prof isn't teaching during this time course time
			if val, ok := teachingMap[p][d]; ok {
				fmt.Println(p, "can't teach", c , "at", course.Assignment.BeginTime ,"since they are already teaching", val, "at", d)
				continue
			}
			// make sure prof isn't teaching too many courses
			max = profsMap[p][c]
			prof = p
		}

		if(max == 195) {
			return prof
		}
	}
	return prof
}

/*
	Input: 	historical-data []Course, SemesterSchedule []Course, professors []Professor
	Output: SemesterSchedule
*/
func AssignCourseProf(historic []structs.Course, semesterSchedule []structs.Course, professors []structs.Professor) []structs.Course {
	
	// get list profs and list of prof preferences
	profsMap, profList := mapPreferences(professors)
	// teachingMap[prof][MthstartTime or TWFstartTime] = courseTitle
	var teachingMap = map[string]map[string]string{}
	
	// for loop through courses needed to be assigned this semester and assign each of them profs
	for i, c := range semesterSchedule {
		prof := assignProf(profsMap, profList, teachingMap, c)

		var t = c.Assignment.BeginTime
		var d = "TWF"+t
		if(c.Assignment.Monday == true){
			d = "MTh"+t
		}		
		
		teachingMap[prof] = map[string]string{}
		teachingMap[prof][d] = c.CourseTitle
		semesterSchedule[i].Prof.DisplayName = prof
	}

	return semesterSchedule
}

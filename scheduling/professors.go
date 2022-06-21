package scheduling

import (
	"algorithm-1/structs"
	"math/rand"
	"time"
)

var j = 0

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
func MapPreferences(profs []structs.Professor) (map[string]map[string]int, []string){
	var profsMap = map[string]map[string]int{}
	var profList []string

	for _, s := range profs {
		profsMap[s.DisplayName] = map[string]int{}
		profList = append(profList, s.DisplayName)
		for _, x := range s.Preferences {
			profsMap[s.DisplayName][x.CourseNum] = int(x.PreferenceNum)
		}
	}
	// return profsMap, profList
	return profsMap, randomizer(profList)
}

/*
	Input:  profsMap map[string]map[string]int, profList []string, teachingMap map[string]map[string]string, course string
	Output: prof string
*/
func assignProf(profsMap map[string]map[string]int, profList []string, teachingMap map[string]string, course structs.Course) (string){
	var max int = 0
	var prof string = "N/A"

	var d string
	if(course.Assignment.Monday == true){
		d = "MTh"+course.Assignment.BeginTime
	} else {
		d = "TWF"+course.Assignment.BeginTime
	}
		
	var c = course.Subject+course.CourseNumber
	var size = len(profList)

	for i := 0; i < size; i++ {
		p := profList[j]

		// make sure prof isn't teaching during this time course time
		if _, skip := teachingMap[p+d]; skip {
			j = (j + 1) % size
			continue
		}

		if max < profsMap[p][c]{
			max = profsMap[p][c]
			prof = profList[j]
		}

		if(max == 7){
			j = (j + 1) % size
			return prof
		}
		j = (j + 1) % size
	}

	return prof
}

/*
	Input: 	historical-data []Course, SemesterSchedule []Course, professors []Professor
	Output: SemesterSchedule
*/
func AssignCourseProf(historic []structs.Course, semesterSchedule []structs.Course, professors []structs.Professor) []structs.Course {
	
	// get list profs and list of prof preferences
	profsMap, profList := MapPreferences(professors)
	var teachingMap = map[string]string{}
	
	// for loop through courses needed to be assigned this semester and assign each of them profs
	for i, c := range semesterSchedule {
		prof := assignProf(profsMap, profList, teachingMap, c)
		
		var d string
		if(c.Assignment.Monday == true){
			d = "MTh"+c.Assignment.BeginTime
		} else {
			d = "TWF"+c.Assignment.BeginTime
		}
		
		teachingMap[prof+d] = c.CourseTitle
		semesterSchedule[i].Prof.DisplayName = prof
	}

	return semesterSchedule
}

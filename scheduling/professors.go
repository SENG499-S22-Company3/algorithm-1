package scheduling

import (
	"algorithm-1/structs"
)

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
	return profsMap, profList
}

/*
	Input:  profsMap map[string]map[string]int, profList []string, course string
	Output: prof string
*/
func assignProf(profsMap map[string]map[string]int, profList []string, course string) (string){
	var max int = 0
	var prof string = "N/A"
	
	for _, p := range profList {
		if max < profsMap[p][course]{
			// make sure prof isn't teaching during this time course time
			// make sure prof isn't teaching too many courses
			max = profsMap[p][course]
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
	
	// for loop through courses needed to be assigned this semester and assign each of them profs
	for i, c := range semesterSchedule {
		prof := assignProf(profsMap, profList, c.Subject+c.CourseNumber)
		semesterSchedule[i].Prof.DisplayName = prof
	}
	
	return semesterSchedule
}

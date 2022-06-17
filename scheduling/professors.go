package scheduling

import (
	"algorithm-1/structs"
	"io/ioutil"
	"log"
)

// "professors":[
//     {
//         "displayName": "Berg, Celina",
//         "prefs": {
// 				"CSC111": 78,
// 				"CSC115": 20,
// 				...
// 			}
//     },
// 	...
// ]

// type ProfMap struct {
// 	DisplayName map[string]PrefMap
// }

// type PrefMap struct {
// 	Preference 	map[string]int
// }

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
Input: (profList, profsMap, course string)
Output: (prof string)
*/
func assignProf(profsMap map[string]map[string]int, profList []string, course structs.Course) (string){
	var max int = 0
	var prof string = "N/A"
	for _, s := range profList {
		if max < profsMap[s][course.CourseTitle]{

			// make sure prof isn't teaching during this time course time
			// make sure prof isn't teaching too many courses

			max = profsMap[s][course.CourseTitle]
			prof = s
		}

		if(max == 195) {
			return prof
		}
	}
	return prof
}

/*
Input: (historical-data, SemesterSchedule, professors)
Output: (SemesterSchedule)
*/
func AssignCourseProf() structs.Schedule {
	
	var schedule structs.Schedule
	jsonData, err := ioutil.ReadFile("./tests/data/input-test.json")
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
	input, _ := structs.ParseInput(jsonData)
	profsMap, profList := mapPreferences(input.Professors)
	
	// for loop through courses needed to be assigned this semester
	// and assign each of them profs
	for {
		prof := assignProf(profsMap, profList, course)
		// add prof to schedule
	}
	
	return schedule
}

package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"math/rand"
	"time"
)

func randomizer(profList []string) []string {
	rand.Seed(time.Now().UnixMilli())
	for i := range profList {
		k := rand.Intn(i + 1)
		profList[i], profList[k] = profList[k], profList[i]
	}
	return profList
}

/*
	Input:  profs []structs.Professor
	Output: prefsMap map[string]map[string]int, 
			profList []string, 
			teachingPrefMax map[string]int
*/
func MapPreferences(profs []structs.Professor, term string) (map[string]map[string]int, []string,  map[string]int){
	var prefsMap = map[string]map[string]int{}
	var profList []string
	var teachingPrefMax = map[string]int{}

	for _, s := range profs {
		
		// set profList
		profList = append(profList, s.DisplayName)

		// set prof preference map
		prefsMap[s.DisplayName] = map[string]int{}
		for _, x := range s.Preferences {
			prefsMap[s.DisplayName][x.CourseNum] = int(x.PreferenceNum)
		}

		// set max prefered courses to teach
		if term == "Fall" {
			teachingPrefMax[s.DisplayName] = int(s.FallTermCourses)
		} else if term == "Spring" {
			teachingPrefMax[s.DisplayName] = int(s.SpringTermCourses)
		} else {
			teachingPrefMax[s.DisplayName] = int(s.SummerTermCourses)
		}
		
	}
	return prefsMap, randomizer(profList), teachingPrefMax
}

/*
	Input:  prefsMap map[string]map[string]int, 
			profList []string, 
			teachingTimeslotMap map[string]map[string]string, 
			teachingCount map[string]int, 
			teachingPrefMax map[string]int,
			course structs.Course, 
			course string
	Output: prof string
*/
func assignProf(prefsMap map[string]map[string]int, 
				profList []string, 
				teachingTimeslotMap map[string]string, 
				teachingCount map[string]int, 
				teachingPrefMax map[string]int,  
				course structs.Course, 
				profPos int) (string, int){
	var max int = 0
	var prof string = "TBD"

	var d string
	if course.Assignment.Monday {
		d = "MTh" + course.Assignment.BeginTime
	} else {
		d = "TWF" + course.Assignment.BeginTime
	}

	var c = course.Subject + course.CourseNumber
	var size = len(profList)

	for i := 0; i < size; i++ {
		// get professor at index profPos
		p := profList[profPos]

		// make sure prof isn't teaching during this course time
		if _, exists := teachingTimeslotMap[p+d]; exists {
			profPos = (profPos + 1) % size
			continue
		}

		// check if profs preference is higher then current
		// and that profs aren't assigned more then there prefered max 
		if max < prefsMap[p][c] && teachingCount[p] < teachingPrefMax[p]{
			max = prefsMap[p][c]
			prof = profList[profPos]
		}

		// if prof has max preference return prof and profPos
		if max == 6 {
			profPos = (profPos + 1) % size
			return prof, profPos
		}
		profPos = (profPos + 1) % size
	}

	return prof, profPos
}

/*
	Input: 	historical-data []Course, 
			SemesterSchedule []Course, 
			professors []Professor,
			term string
	Output: SemesterSchedule
*/
func AssignCourseProf(historic []structs.Course, semesterSchedule []structs.Course, professors []structs.Professor, term string) []structs.Course {
	
	// get list profs and list of prof preferences
	prefsMap, profList, teachingPrefMax := MapPreferences(professors, term)
	var teachingCount = map[string]int{}
	var teachingTimeslotMap = map[string]string{}
	var courseMap = map[string]string{}
	var prof string
	var profPos = 0
	var d string

	// for loop through courses needed to be assigned this semester and assign each of them profs
	for i, c := range semesterSchedule {

		// need to check if professors has taught more then prefered courses this semester
		if val, exists := courseMap[c.Subject+c.CourseNumber]; exists && teachingCount[val] < teachingPrefMax[val] {
			prof = val
		}else{
			prof, profPos = assignProf(prefsMap, profList, teachingTimeslotMap, teachingCount, teachingPrefMax, c, profPos)
		}
		
		if prof != "TBA" {
			if(c.Assignment.Monday == true){
				d = "MTh"+c.Assignment.BeginTime
			} else {
				d = "TWF"+c.Assignment.BeginTime
			}
			// update map used to asssign same prof to different sections of the same course
			courseMap[c.Subject+c.CourseNumber] = prof  	
			// update map used to ensure teachers aren't double slotted 
			teachingTimeslotMap[prof+d] = c.CourseTitle	
			// increase prof teaching count
			teachingCount[prof]++
		}
				
		// update semester schedule
		semesterSchedule[i].Prof.DisplayName = prof
	}

	return semesterSchedule
}

func ScheduleConstraintsCheck(term string, testScheduleCourse []structs.Course, input structs.Input) error {

	var teachingMap = map[string]map[string]string{}
	var d string
	var err error

	for _, p := range input.Professors {
		teachingMap[p.DisplayName] = map[string]string{}
	}

	prefsMap, _, _ := MapPreferences(input.Professors, term)

	for _, c := range testScheduleCourse {
		if c.Assignment.Monday {
			d = "MTh" + c.Assignment.BeginTime
		} else {
			d = "TWF" + c.Assignment.BeginTime
		}

		if c.Prof.DisplayName == "" || c.Assignment.BeginTime == "" || c.Assignment.EndTime == "" {
			err = fmt.Errorf("error: %v Schedule missing %v %v timeslot and/or prof,   ", term, c.Subject, c.CourseNumber)
			break
		}

		if _, found := teachingMap[c.Prof.DisplayName][d]; found {
			err = fmt.Errorf("error: %v teaching another %v course at %v,   ", c.Prof.DisplayName, term, d)
			break
		}

		if val, pass := prefsMap[c.Prof.DisplayName][c.Subject+c.CourseNumber]; !pass && c.Prof.DisplayName != "TBD" {
			err = fmt.Errorf(c.Prof.DisplayName, "cannot teach this "+term+" course since they have no (", val, ") preference,   ")
			break
		}

		teachingMap[c.Prof.DisplayName][d] = c.CourseTitle + d

		// need to add prefered number of courses to teach a semester logic once this feature
		// is put in production
	}

	return err
}

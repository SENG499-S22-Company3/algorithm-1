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

func countHardScheduled(hardScheduled []structs.Course) map[string]int {
	var courseMap = map[string]int{}
	for _, c := range hardScheduled {
		if c.Prof.DisplayName != "TBD" {
			courseMap[c.Prof.DisplayName]++
		}
	}
	return courseMap
}

/*
	Input:  prefsMap map[string]map[string]int, 
			profList []string, 
			teachingTimeslotMap map[string]map[string]string, 
			teachingCount map[string]int, 
			teachingPrefMax map[string]int,
			course structs.Course,
			profPos int
	Output: prof string, profPos int
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

		if teachingCount[p] < teachingPrefMax[p] {
			// make sure prof isn't teaching during this course time
			if _, exists := teachingTimeslotMap[p+d]; exists {
				profPos = (profPos + 1) % size
				continue
			}

			// check if profs preference is higher then current
			// and that profs aren't assigned more then there prefered max 
			if max < prefsMap[p][c] {
				max = prefsMap[p][c]
				prof = profList[profPos]
			}

			// if prof has max preference return prof and profPos
			if max == 6 {
				profPos = (profPos + 1) % size
				return prof, profPos
			}
		}
		profPos = (profPos + 1) % size
	}

	return prof, profPos
}

/*
	Input: 	hardScheduled []Course, 
			SemesterSchedule []Course, 
			professors []Professor,
			term string
	Output: SemesterSchedule
*/
func AssignCourseProf(hardScheduled []structs.Course, semesterSchedule []structs.Course, professors []structs.Professor, term string) []structs.Course {
	
	// get list profs and list of prof preferences
	prefsMap, profList, teachingPrefMax := MapPreferences(professors, term)
	teachingCount := countHardScheduled(hardScheduled)
	var teachingTimeslotMap = map[string]string{}
	var courseMap  = map[string]string{}
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
		semesterSchedule[i].Prof.Preferences = make([]structs.Preference, 0)
	}

	return semesterSchedule
}

func ScheduleConstraintsCheck(term string,
	testScheduleCourse []structs.Course,
	profs []structs.Professor) error {

	var teachingTimeslotMap = map[string]string{}
	var teachingCount = map[string]int{}
	var courseMap = map[string]string{}
	var d string
	var err error

	prefsMap, profList, teachingPrefMax := MapPreferences(profs, term)

	for _, c := range testScheduleCourse {
		if c.Assignment.Monday {
			d = "MTh" + c.Assignment.BeginTime
		} else {
			d = "TWF" + c.Assignment.BeginTime
		}

		// check for unscheduled course
		if c.Prof.DisplayName == "" || c.Assignment.BeginTime == "" || c.Assignment.EndTime == "" {
			err = fmt.Errorf("error: %v Schedule missing %v %v timeslot and/or prof,   ", term, c.Subject, c.CourseNumber)
			break
		}

		// check for double slotted prof
		if _, found := teachingTimeslotMap[c.Prof.DisplayName+d]; found {
			err = fmt.Errorf("error: %v teaching another %v course at %v,   ", c.Prof.DisplayName, term, d)
			break
		}

		// check that prof with zero preference doesn't get sheduled
		if pref, pass := prefsMap[c.Prof.DisplayName][c.Subject+c.CourseNumber]; !pass && c.Prof.DisplayName != "TBD" {
			err = fmt.Errorf(c.Prof.DisplayName, "cannot teach this "+term+" course since they have no (", pref, ") preference,   ")
			break
		}

		// check that while a prof is teaching less than the number of prefered courese they will teach all sections
		if prof, exists := courseMap[c.Subject+c.CourseNumber]; exists && !(teachingCount[prof] < teachingPrefMax[prof]) && c.Prof.DisplayName == prof {
			fmt.Println(teachingCount[prof] , teachingPrefMax[prof])
			err = fmt.Errorf("error: %v should be teaching this %v since they teach other sections.", prof, c.Subject+c.CourseNumber)
			break
		}

		if c.Prof.DisplayName != "TBD" {
			// update map used to ensure teachers aren't double slotted 
			teachingTimeslotMap[c.Prof.DisplayName+d] = c.CourseTitle + d
			// update map used to asssign same prof to different sections of the same course
			courseMap[c.Subject+c.CourseNumber] = c.Prof.DisplayName 	
			// increase prof teaching count
			teachingCount[c.Prof.DisplayName]++
		}
	}

	for _, prof := range profList {
		// check that profs do not teach more than prefered amount of courses
		if teachingCount[prof] > teachingPrefMax[prof] {
			err = fmt.Errorf("error: %v teaching %v courses which is more than there prefered %v courses.", prof, teachingCount[prof], teachingPrefMax[prof])
			break
		}
	}

	return err
}

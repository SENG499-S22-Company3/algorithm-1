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
	Output: prefsMap map[string]map[string]int, profList []string
*/
func MapPreferences(profs []structs.Professor) (map[string]map[string]int, []string) {
	var prefsMap = map[string]map[string]int{}
	var profList []string

	for _, s := range profs {
		prefsMap[s.DisplayName] = map[string]int{}
		profList = append(profList, s.DisplayName)
		for _, x := range s.Preferences {
			prefsMap[s.DisplayName][x.CourseNum] = int(x.PreferenceNum)
		}
	}
	return prefsMap, randomizer(profList)
}

/*
	Input:  prefsMap map[string]map[string]int, profList []string, teachingMap map[string]string, course string
	Output: prof string
*/
func assignProf(prefsMap map[string]map[string]int, profList []string, teachingMap map[string]string, course structs.Course, profPos int) (string, int) {
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

		// make sure prof isn't teaching during this time course time
		if _, skip := teachingMap[p+d]; skip {
			profPos = (profPos + 1) % size
			continue
		}

		// check if profs preference is higher then current
		if max < prefsMap[p][c] {
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
	Input: 	historical-data []Course, SemesterSchedule []Course, professors []Professor
	Output: SemesterSchedule
*/
func AssignCourseProf(historic []structs.Course, semesterSchedule []structs.Course, professors []structs.Professor) []structs.Course {

	// get list profs and list of prof preferences
	prefsMap, profList := MapPreferences(professors)
	var teachingMap = map[string]string{}
	var courseMap = map[string]string{}
	var prof string
	var profPos = 0
	var d string

	// for loop through courses needed to be assigned this semester and assign each of them profs
	for i, c := range semesterSchedule {

		// need to check if professors has taught more then prefered courses this semester
		if val, skip := courseMap[c.Subject+c.CourseNumber]; skip {
			prof = val
			// TODO: increase prof teaching count
		} else {
			prof, profPos = assignProf(prefsMap, profList, teachingMap, c, profPos)
		}

		if c.Assignment.Monday {
			d = "MTh" + c.Assignment.BeginTime
		} else {
			d = "TWF" + c.Assignment.BeginTime
		}

		// update map used to asssign same prof to different sections of the same course
		courseMap[c.Subject+c.CourseNumber] = prof
		// update map used to ensure teachers aren't double slotted
		teachingMap[prof+d] = c.CourseTitle
		// update semester schedule
		semesterSchedule[i].Prof.DisplayName = prof
		semesterSchedule[i].Prof.Preferences = []structs.Preference{}
	}

	return semesterSchedule
}

func ScheduleConstraintsCheck(term string,
	testScheduleCourse []structs.Course,
	profs []structs.Professor) error {

	var teachingMap = map[string]string{}
	var d string
	var err error

	prefsMap, _ := MapPreferences(profs)

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

		if _, found := teachingMap[c.Prof.DisplayName+d]; found {
			err = fmt.Errorf("error: %v teaching another %v course at %v,   ", c.Prof.DisplayName, term, d)
			break
		}

		if val, pass := prefsMap[c.Prof.DisplayName][c.Subject+c.CourseNumber]; !pass && c.Prof.DisplayName != "TBD" {
			err = fmt.Errorf(c.Prof.DisplayName, "cannot teach this "+term+" course since they have no (", val, ") preference,   ")
			break
		}

		if c.Prof.DisplayName != "TBD" {
			teachingMap[c.Prof.DisplayName+d] = c.CourseTitle + d
		}
	}

	return err
}

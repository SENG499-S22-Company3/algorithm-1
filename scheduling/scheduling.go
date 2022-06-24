package scheduling

import (
	"algorithm-1/structs"
)

func sliceContains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func BaseSemester(requestedCourses []structs.Course, historicalSemester []structs.Course) []structs.Course {
	var result []structs.Course
	ignore := []string{"CHEM", "MATH", "PHYS", "STAT", "ECON"} // Not sure if this is comprehensive
	for _, h := range historicalSemester {
		if sliceContains(h.Subject, ignore) {
			result = append(result, h)
		}
		for _, c := range requestedCourses {
			if c.Subject == h.Subject && c.CourseNumber == h.CourseNumber {
				h.StreamSequence = c.StreamSequence
				result = append(result, h)
			}
		}
	}
	return result
}

func BaseSchedule(requestedCourses structs.Schedule, historicalSchedule structs.Schedule) structs.Schedule {
	// making channels to get return values from goroutines
	fall := make(chan []structs.Course)
	spring := make(chan []structs.Course)
	summer := make(chan []structs.Course)

	go func() {
		fall <- BaseSemester(requestedCourses.FallCourses, historicalSchedule.FallCourses)
	}()
	go func() {
		spring <- BaseSemester(requestedCourses.SpringCourses, historicalSchedule.SpringCourses)
	}()
	go func() {
		summer <- BaseSemester(requestedCourses.SummerCourses, historicalSchedule.SummerCourses)
	}()

	return structs.Schedule{
		FallCourses:   <-fall,
		SpringCourses: <-spring,
		SummerCourses: <-summer,
	}
}

func Assignments(historicalSemester []structs.Course, requestedCourses []structs.Course, professors []structs.Professor) ([]structs.Course){
	
	timeslotFallMap, _ := BaseTimeslotMaps(historicalSemester)


	requestedCourses, _, _ = AddCoursesToStreamMaps(requestedCourses, timeslotFallMap)

	requestedCourses = AssignCourseProf(historicalSemester, requestedCourses, professors)

	return requestedCourses
}

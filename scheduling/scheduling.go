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

func BaseSchedule(requestedCourses []structs.Course, historicalSymester []structs.Course) []structs.Course {
	var result []structs.Course
	ignore := []string{"CHEM", "MATH", "PHYS", "STAT", "ECON"} // Not sure if this is comprehensive
	for _, h := range historicalSymester {
		if sliceContains(h.Subject, ignore) {
			result = append(result, h)
		}
		for _, c := range requestedCourses {
			if c.Subject == h.Subject && c.CourseNumber == h.CourseNumber {
				result = append(result, h)
			}
		}
	}
	return result
}

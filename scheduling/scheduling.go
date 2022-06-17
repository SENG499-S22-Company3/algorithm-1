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

func CreateEmptyStreamType() structs.StreamType {
	emptyMTh := map[string]string{
		"0830": "",
		"1000": "",
		"1130": "",
		"1300": "",
		"1430": "",
		"1600": "",
		"1730": "",
	}

	emptyTWF := map[string]string{
		"0830": "",
		"0930": "",
		"1030": "",
		"1130": "",
		"1230": "",
		"1330": "",
		"1430": "",
		"1530": "",
		"1630": "",
		"1730": "",
	}

	emptyTimeslots := structs.Timeslots{
		Monday:    emptyMTh,
		Tuesday:   emptyTWF,
		Wednesday: emptyTWF,
		Thursday:  emptyMTh,
		Friday:    emptyTWF,
	}

	timeslotMaps := structs.StreamType{
		S1A: emptyTimeslots,
		S1B: emptyTimeslots,
		S2A: emptyTimeslots,
		S2B: emptyTimeslots,
		S3A: emptyTimeslots,
		S4A: emptyTimeslots,
		S4B: emptyTimeslots,
	}

	return timeslotMaps
}

func BaseTimeslotMaps(historicTermCourses []structs.Course) structs.StreamType {
	timeslotMaps := CreateEmptyStreamType()

	return timeslotMaps
}

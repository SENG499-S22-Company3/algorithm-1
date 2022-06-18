package scheduling

import (
	"algorithm-1/structs"
	"fmt"
)

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
		S3B: emptyTimeslots,
		S4A: emptyTimeslots,
		S4B: emptyTimeslots,
	}

	return timeslotMaps
}

func BaseTimeslotMaps(baseTermCourses []structs.Course) structs.StreamType {
	timeslotMaps := CreateEmptyStreamType()

	for _, course := range baseTermCourses {
		if course.StreamSequence == "1A" {
			AddMultipleTimeslots(course, timeslotMaps.S1A)
		} else if course.StreamSequence == "1B" {
			AddMultipleTimeslots(course, timeslotMaps.S1B)
		} else if course.StreamSequence == "2A" {
			AddMultipleTimeslots(course, timeslotMaps.S2A)
		} else if course.StreamSequence == "2B" {
			AddMultipleTimeslots(course, timeslotMaps.S2B)
		} else if course.StreamSequence == "3A" {
			AddMultipleTimeslots(course, timeslotMaps.S3A)
		} else if course.StreamSequence == "3B" {
			AddMultipleTimeslots(course, timeslotMaps.S3B)
		} else if course.StreamSequence == "4A" {
			AddMultipleTimeslots(course, timeslotMaps.S4A)
		} else if course.StreamSequence == "4B" {
			AddMultipleTimeslots(course, timeslotMaps.S4B)
		}
	}

	return timeslotMaps
}

func AddMultipleTimeslots(course structs.Course, timeslots structs.Timeslots) {

	if course.Assignment.BeginTime != "" {
		if course.Assignment.Monday {
			AddTimeslot(course, timeslots.Monday)
		}
		if course.Assignment.Tuesday {
			AddTimeslot(course, timeslots.Tuesday)
		}
		if course.Assignment.Wednesday {
			AddTimeslot(course, timeslots.Wednesday)
		}
		if course.Assignment.Thursday {
			AddTimeslot(course, timeslots.Thursday)
		}
		if course.Assignment.Friday {
			AddTimeslot(course, timeslots.Friday)
		}
	} else {
		// TO DO Handle non-historic courses
	}
}

func AddTimeslot(course structs.Course, day map[string]string) {

	if _, isValid := day[course.Assignment.BeginTime]; !isValid {
		fmt.Printf("Error: %v %v is scheduled outside of valid lecture time", course.Subject, course.CourseNumber)
	} else if scheduledCourse := day[course.Assignment.BeginTime]; scheduledCourse != "" {
		fmt.Printf("Error: %v %v is scheduled at same time as another required course %v", course.Subject, course.CourseNumber, scheduledCourse)
	} else {
		day[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
	}
}

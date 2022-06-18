package scheduling

import (
	"algorithm-1/structs"
	"fmt"
)

func CreateEmptyStreamType() structs.StreamType {

	timeslotMaps := structs.StreamType{
		S1A: structs.Timeslots{
			Monday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Tuesday: map[string]string{
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
			},
			Wednesday: map[string]string{
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
			},
			Thursday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Friday: map[string]string{
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
			},
		},
		S1B: structs.Timeslots{
			Monday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Tuesday: map[string]string{
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
			},
			Wednesday: map[string]string{
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
			},
			Thursday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Friday: map[string]string{
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
			},
		},
		S2A: structs.Timeslots{
			Monday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Tuesday: map[string]string{
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
			},
			Wednesday: map[string]string{
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
			},
			Thursday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Friday: map[string]string{
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
			},
		},
		S2B: structs.Timeslots{
			Monday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Tuesday: map[string]string{
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
			},
			Wednesday: map[string]string{
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
			},
			Thursday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Friday: map[string]string{
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
			},
		},
		S3A: structs.Timeslots{
			Monday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Tuesday: map[string]string{
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
			},
			Wednesday: map[string]string{
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
			},
			Thursday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Friday: map[string]string{
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
			},
		},
		S3B: structs.Timeslots{
			Monday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Tuesday: map[string]string{
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
			},
			Wednesday: map[string]string{
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
			},
			Thursday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Friday: map[string]string{
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
			},
		},
		S4A: structs.Timeslots{
			Monday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Tuesday: map[string]string{
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
			},
			Wednesday: map[string]string{
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
			},
			Thursday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Friday: map[string]string{
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
			},
		},
		S4B: structs.Timeslots{
			Monday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Tuesday: map[string]string{
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
			},
			Wednesday: map[string]string{
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
			},
			Thursday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
				"1300": "",
				"1430": "",
				"1600": "",
				"1730": "",
			},
			Friday: map[string]string{
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
			},
		},
	}

	return timeslotMaps
}

func BaseTimeslotMaps(baseTermCourses []structs.Course) (structs.StreamType, string) {
	timeslotMaps := CreateEmptyStreamType()
	err := ""

	for _, course := range baseTermCourses {
		if course.StreamSequence == "1A" {
			timeslotMaps.S1A, err = AddMultipleTimeslots(course, timeslotMaps.S1A)
		} else if course.StreamSequence == "1B" {
			timeslotMaps.S1B, err = AddMultipleTimeslots(course, timeslotMaps.S1B)
		} else if course.StreamSequence == "2A" {
			timeslotMaps.S2A, err = AddMultipleTimeslots(course, timeslotMaps.S2A)
		} else if course.StreamSequence == "2B" {
			timeslotMaps.S2B, err = AddMultipleTimeslots(course, timeslotMaps.S2B)
		} else if course.StreamSequence == "3A" {
			timeslotMaps.S3A, err = AddMultipleTimeslots(course, timeslotMaps.S3A)
		} else if course.StreamSequence == "3B" {
			timeslotMaps.S3B, err = AddMultipleTimeslots(course, timeslotMaps.S3B)
		} else if course.StreamSequence == "4A" {
			timeslotMaps.S4A, err = AddMultipleTimeslots(course, timeslotMaps.S4A)
		} else if course.StreamSequence == "4B" {
			timeslotMaps.S4B, err = AddMultipleTimeslots(course, timeslotMaps.S4B)
		}
	}

	return timeslotMaps, err
}

func AddMultipleTimeslots(course structs.Course, timeslots structs.Timeslots) (structs.Timeslots, string) {
	err := ""

	if course.Assignment.BeginTime != "" {
		if course.Assignment.Monday {
			timeslots.Monday, err = AddTimeslot(course, timeslots.Monday)
		}
		if course.Assignment.Tuesday {
			timeslots.Tuesday, err = AddTimeslot(course, timeslots.Tuesday)
		}
		if course.Assignment.Wednesday {
			timeslots.Wednesday, err = AddTimeslot(course, timeslots.Wednesday)
		}
		if course.Assignment.Thursday {
			timeslots.Thursday, err = AddTimeslot(course, timeslots.Thursday)
		}
		if course.Assignment.Friday {
			timeslots.Friday, err = AddTimeslot(course, timeslots.Friday)
		}
	} else {
		// TO DO Handle non-historic courses
	}

	return timeslots, err
}

func AddTimeslot(course structs.Course, day map[string]string) (map[string]string, string) {
	err := ""

	if _, isValid := day[course.Assignment.BeginTime]; !isValid {
		err = fmt.Sprintf("Error: %v %v is scheduled outside of valid lecture time", course.Subject, course.CourseNumber)
	} else if scheduledCourse := day[course.Assignment.BeginTime]; scheduledCourse != "" {
		err = fmt.Sprintf("Error: %v %v is scheduled at same time as another required course %v", course.Subject, course.CourseNumber, scheduledCourse)
	} else {
		day[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
	}

	return day, err
}

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

	err = AddCoursesToStreamMaps(baseTermCourses, timeslotMaps)

	return timeslotMaps, err
}

func AddCoursesToStreamMaps(courses []structs.Course, timeslotMaps structs.StreamType) string {
	err := ""

	for _, course := range courses {
		if course.StreamSequence == "1A" {
			err = AddMultipleTimeslots(course, timeslotMaps.S1A)
		} else if course.StreamSequence == "1B" {
			err = AddMultipleTimeslots(course, timeslotMaps.S1B)
		} else if course.StreamSequence == "2A" {
			err = AddMultipleTimeslots(course, timeslotMaps.S2A)
		} else if course.StreamSequence == "2B" {
			err = AddMultipleTimeslots(course, timeslotMaps.S2B)
		} else if course.StreamSequence == "3A" {
			err = AddMultipleTimeslots(course, timeslotMaps.S3A)
		} else if course.StreamSequence == "3B" {
			err = AddMultipleTimeslots(course, timeslotMaps.S3B)
		} else if course.StreamSequence == "4A" {
			err = AddMultipleTimeslots(course, timeslotMaps.S4A)
		} else if course.StreamSequence == "4B" {
			err = AddMultipleTimeslots(course, timeslotMaps.S4B)
		}
	}

	return err
}

func AddMultipleTimeslots(course structs.Course, timeslots structs.Timeslots) string {
	err := ""

	if course.Assignment.BeginTime != "" {
		if course.Assignment.Monday {
			err = AddTimeslot(course, timeslots.Monday)
		}
		if course.Assignment.Tuesday {
			err = AddTimeslot(course, timeslots.Tuesday)
		}
		if course.Assignment.Wednesday {
			err = AddTimeslot(course, timeslots.Wednesday)
		}
		if course.Assignment.Thursday {
			err = AddTimeslot(course, timeslots.Thursday)
		}
		if course.Assignment.Friday {
			err = AddTimeslot(course, timeslots.Friday)
		}
	} else {
		// TO DO Handle non-historic courses
	}

	return err
}

func AddTimeslot(course structs.Course, day map[string]string) string {
	err := ""

	if _, isValid := day[course.Assignment.BeginTime]; !isValid {
		err = fmt.Sprintf("Error: %v %v is scheduled during a regular block time at %v", course.Subject, course.CourseNumber, course.Assignment.BeginTime)
	} else if scheduledCourse := day[course.Assignment.BeginTime]; scheduledCourse != "" {
		err = fmt.Sprintf("Error: %v %v is scheduled at same time as another required course %v", course.Subject, course.CourseNumber, scheduledCourse)
	} else {
		day[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
	}

	return err
}

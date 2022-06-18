package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"math/rand"
	"strconv"
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
	hasBeenAdded := false

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
		for !hasBeenAdded {
			selection := rand.Intn(2) // Create random integer to decide whether or not to choose MTh, or TWF

			if selection == 0 {
				for time, courseValue := range timeslots.Monday {
					if courseValue == "" {
						if timeslots.Thursday[time] == "" {
							course = SetCourseTime(course, time, true)
							AddTimeslot(course, timeslots.Monday)
							AddTimeslot(course, timeslots.Thursday)
							hasBeenAdded = true
						}
					}
				}
			} else {
				for time, courseValue := range timeslots.Tuesday {
					if courseValue == "" {
						if timeslots.Wednesday[time] == "" && timeslots.Friday[time] == "" {
							course = SetCourseTime(course, time, false)
							AddTimeslot(course, timeslots.Tuesday)
							AddTimeslot(course, timeslots.Wednesday)
							AddTimeslot(course, timeslots.Friday)
							hasBeenAdded = true
						}
					}
				}
			}
		}
	}

	return err
}

func AddTimeslot(course structs.Course, day map[string]string) string {
	err := ""

	if _, isValid := day[course.Assignment.BeginTime]; !isValid { // Check if map key exists
		err = fmt.Sprintf("Error: %v %v is scheduled during a regular block time at %v", course.Subject, course.CourseNumber, course.Assignment.BeginTime)
	} else if scheduledCourse := day[course.Assignment.BeginTime]; scheduledCourse != "" { // Check if there is already a course there
		err = fmt.Sprintf("Error: %v %v is scheduled at same time as another required course %v", course.Subject, course.CourseNumber, scheduledCourse)
	} else {
		day[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
	}

	return err
}

func SetCourseTime(course structs.Course, beginTime string, isMTh bool) structs.Course {
	course.Assignment.BeginTime = beginTime
	beginTimeInt, _ := strconv.Atoi(beginTime)

	if isMTh {
		course.Assignment.EndTime = strconv.Itoa(beginTimeInt + 120)
		course.Assignment.Monday = true
		course.Assignment.Thursday = true
	} else {
		course.Assignment.EndTime = strconv.Itoa(beginTimeInt + 90)
		course.Assignment.Tuesday = true
		course.Assignment.Wednesday = true
		course.Assignment.Friday = true
	}

	if len(course.Assignment.EndTime) == 3 {
		course.Assignment.EndTime = "0" + course.Assignment.EndTime
	}

	return course
}

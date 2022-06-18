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

func BaseTimeslotMaps(baseTermCourses []structs.Course) (structs.StreamType, error) {
	timeslotMaps := CreateEmptyStreamType()

	_, timeslotMaps, err := AddCoursesToStreamMaps(baseTermCourses, timeslotMaps)

	return timeslotMaps, err
}

func AddCoursesToStreamMaps(courses []structs.Course, timeslotMaps structs.StreamType) ([]structs.Course, structs.StreamType, error) {
	var err error
	var updatedCourses []structs.Course
	var updatedCourse structs.Course

	for _, course := range courses {
		if course.StreamSequence == "1A" {
			updatedCourse, timeslotMaps.S1A, err = AddMultipleTimeslots(course, timeslotMaps.S1A)
		} else if course.StreamSequence == "1B" {
			updatedCourse, timeslotMaps.S1B, err = AddMultipleTimeslots(course, timeslotMaps.S1B)
		} else if course.StreamSequence == "2A" {
			updatedCourse, timeslotMaps.S2A, err = AddMultipleTimeslots(course, timeslotMaps.S2A)
		} else if course.StreamSequence == "2B" {
			updatedCourse, timeslotMaps.S2B, err = AddMultipleTimeslots(course, timeslotMaps.S2B)
		} else if course.StreamSequence == "3A" {
			updatedCourse, timeslotMaps.S3A, err = AddMultipleTimeslots(course, timeslotMaps.S3A)
		} else if course.StreamSequence == "3B" {
			updatedCourse, timeslotMaps.S3B, err = AddMultipleTimeslots(course, timeslotMaps.S3B)
		} else if course.StreamSequence == "4A" {
			updatedCourse, timeslotMaps.S4A, err = AddMultipleTimeslots(course, timeslotMaps.S4A)
		} else if course.StreamSequence == "4B" {
			updatedCourse, timeslotMaps.S4B, err = AddMultipleTimeslots(course, timeslotMaps.S4B)
		}

		updatedCourses = append(updatedCourses, updatedCourse)
	}

	return updatedCourses, timeslotMaps, err
}

func AddMultipleTimeslots(course structs.Course, timeslots structs.Timeslots) (structs.Course, structs.Timeslots, error) {
	var err error
	hasBeenAdded := false

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
		for !hasBeenAdded {
			selection := rand.Intn(2) // Create random integer to decide whether or not to choose MTh, or TWF

			if selection == 0 {
				for time, courseValue := range timeslots.Monday {
					if courseValue == "" {
						if timeslots.Thursday[time] == "" {
							course = SetCourseTime(course, time, true)
							timeslots.Monday, err = AddTimeslot(course, timeslots.Monday)
							timeslots.Thursday, err = AddTimeslot(course, timeslots.Thursday)
							hasBeenAdded = true
							break
						}
					}
				}
			} else {
				for time, courseValue := range timeslots.Tuesday {
					if courseValue == "" {
						if timeslots.Wednesday[time] == "" && timeslots.Friday[time] == "" {
							course = SetCourseTime(course, time, false)
							timeslots.Tuesday, err = AddTimeslot(course, timeslots.Tuesday)
							timeslots.Wednesday, err = AddTimeslot(course, timeslots.Wednesday)
							timeslots.Thursday, err = AddTimeslot(course, timeslots.Friday)
							hasBeenAdded = true
							break
						}
					}
				}
			}
		}
	}

	return course, timeslots, err
}

func AddTimeslot(course structs.Course, day map[string]string) (map[string]string, error) {
	var err error

	if _, isValid := day[course.Assignment.BeginTime]; !isValid { // Check if map key exists
		err = fmt.Errorf("error: %v %v is scheduled during a regular block time at %v", course.Subject, course.CourseNumber, course.Assignment.BeginTime)
	} else if scheduledCourse := day[course.Assignment.BeginTime]; scheduledCourse != "" { // Check if there is already a course there
		err = fmt.Errorf("error: %v %v is scheduled at same time as another required course %v", course.Subject, course.CourseNumber, scheduledCourse)
	} else {
		day[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
	}

	return day, err
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

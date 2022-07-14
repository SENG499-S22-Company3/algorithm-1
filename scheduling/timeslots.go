package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"math"
)

func createEmptyDay(isMTh bool) map[string]string {
	var day map[string]string

	if isMTh {
		day = map[string]string{
			"0830": "",
			"1000": "",
			"1130": "",
			"1300": "",
			"1430": "",
			"1600": "",
			"1730": "",
		}
	} else {
		day = map[string]string{
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
	}

	return day
}

func createEmptySequence() structs.Timeslots {
	return structs.Timeslots{
		Monday:    createEmptyDay(true),
		Tuesday:   createEmptyDay(false),
		Wednesday: createEmptyDay(false),
		Thursday:  createEmptyDay(true),
		Friday:    createEmptyDay(false),
	}
}

func CreateEmptyStreamType() structs.StreamType {

	timeslotMaps := structs.StreamType{
		S1A: createEmptySequence(),
		S1B: createEmptySequence(),
		S2A: createEmptySequence(),
		S2B: createEmptySequence(),
		S3A: createEmptySequence(),
		S3B: createEmptySequence(),
		S4A: createEmptySequence(),
		S4B: createEmptySequence(),
	}

	return timeslotMaps
}

func BaseTimeslotMaps(baseTermCourses []structs.Course, term string) (structs.StreamType, error) {
	timeslotMaps := CreateEmptyStreamType()

	_, timeslotMaps, err := AddCoursesToStreamMaps(baseTermCourses, timeslotMaps, term)

	return timeslotMaps, err
}

func AddCoursesToStreamMaps(courses []structs.Course, timeslotMaps structs.StreamType, term string) ([]structs.Course, structs.StreamType, error) {
	var err error
	var updatedCourses []structs.Course
	var updatedCourse structs.Course

	for _, course := range courses {
		if course.StreamSequence == "1A" {
			updatedCourse, timeslotMaps.S1A, err = addMultipleTimeslots(course, timeslotMaps.S1A, term)
		} else if course.StreamSequence == "1B" {
			updatedCourse, timeslotMaps.S1B, err = addMultipleTimeslots(course, timeslotMaps.S1B, term)
		} else if course.StreamSequence == "2A" {
			updatedCourse, timeslotMaps.S2A, err = addMultipleTimeslots(course, timeslotMaps.S2A, term)
		} else if course.StreamSequence == "2B" {
			updatedCourse, timeslotMaps.S2B, err = addMultipleTimeslots(course, timeslotMaps.S2B, term)
		} else if course.StreamSequence == "3A" {
			updatedCourse, timeslotMaps.S3A, err = addMultipleTimeslots(course, timeslotMaps.S3A, term)
		} else if course.StreamSequence == "3B" {
			updatedCourse, timeslotMaps.S3B, err = addMultipleTimeslots(course, timeslotMaps.S3B, term)
		} else if course.StreamSequence == "4A" {
			updatedCourse, timeslotMaps.S4A, err = addMultipleTimeslots(course, timeslotMaps.S4A, term)
		} else if course.StreamSequence == "4B" {
			updatedCourse, timeslotMaps.S4B, err = addMultipleTimeslots(course, timeslotMaps.S4B, term)
		} else {
			err = fmt.Errorf("error: %v %v has no stream sequence value in %v term", course.Subject, course.CourseNumber, term)
		}

		if err != nil {
			break
		}

		updatedCourses = append(updatedCourses, updatedCourse)
	}

	return updatedCourses, timeslotMaps, err
}

func addMultipleTimeslots(course structs.Course, timeslots structs.Timeslots, term string) (structs.Course, structs.Timeslots, error) {
	var err error
	hasBeenAdded := false
	count := 0

	course = setCourseDates(course, term)

	if course.Assignment.BeginTime != "" {
		if course.Assignment.Monday {
			timeslots.Monday, err = addTimeslot(course, timeslots.Monday, term)
		}
		if course.Assignment.Tuesday {
			timeslots.Tuesday, err = addTimeslot(course, timeslots.Tuesday, term)
		}
		if course.Assignment.Wednesday {
			timeslots.Wednesday, err = addTimeslot(course, timeslots.Wednesday, term)
		}
		if course.Assignment.Thursday {
			timeslots.Thursday, err = addTimeslot(course, timeslots.Thursday, term)
		}
		if course.Assignment.Friday {
			timeslots.Friday, err = addTimeslot(course, timeslots.Friday, term)
		}
	} else {
		for !hasBeenAdded && count < 20 {
			selection := rand.Intn(2) // Create random integer to decide whether or not to choose MTh, or TWF

			if selection == 0 {
				for time, courseValue := range timeslots.Monday {
					if courseValue == "" {
						if timeslots.Thursday[time] == "" {
							course = setCourseTime(course, time, true)
							timeslots.Monday, err = addTimeslot(course, timeslots.Monday, term)
							if err != nil {
								break
							}
							timeslots.Thursday, err = addTimeslot(course, timeslots.Thursday, term)
							hasBeenAdded = true
							break
						}
					}
				}
			} else {
				for time, courseValue := range timeslots.Tuesday {
					if courseValue == "" {
						if timeslots.Wednesday[time] == "" && timeslots.Friday[time] == "" {
							course = setCourseTime(course, time, false)
							timeslots.Tuesday, err = addTimeslot(course, timeslots.Tuesday, term)
							if err != nil {
								break
							}
							timeslots.Wednesday, err = addTimeslot(course, timeslots.Wednesday, term)
							if err != nil {
								break
							}
							timeslots.Friday, err = addTimeslot(course, timeslots.Friday, term)
							hasBeenAdded = true
							break
						}
					}
				}
			}
			count++
		}
	}

	if count >= 20 {
		err = fmt.Errorf("error: Ran out of slots to assign courses in stream %v in the %v term", course.StreamSequence, term)
	}

	return course, timeslots, err
}

func addTimeslot(course structs.Course, day map[string]string, term string) (map[string]string, error) {
	var err error
	//beginTimeInt, _ := strconv.Atoi(course.Assignment.BeginTime)
	//endTimeInt, _ := strconv.Atoi(course.Assignment.EndTime)

	if _, isValid := day[course.Assignment.BeginTime]; !isValid { // Check if map key exists
		err = fmt.Errorf("error: %v %v is scheduled outside a regular block time at %v in  %v term   ", course.Subject, course.CourseNumber, course.Assignment.BeginTime, term)
	} else if scheduledCourse := day[course.Assignment.BeginTime]; scheduledCourse != "" { // Check if there is already a course there
		err = fmt.Errorf("error: %v %v is scheduled at same time as another required course %v in %v term,   ", course.Subject, course.CourseNumber, scheduledCourse, term)
	} else {
		day[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
	}
	/*
		if endTimeInt-beginTimeInt == 250 || endTimeInt-beginTimeInt == 290 { // Check if three hour course
			// TO-DO handle 3 hour courses
		}
	*/

	return day, err
}

func setCourseTime(course structs.Course, beginTime string, isMTh bool) structs.Course {
	course.Assignment.BeginTime = beginTime
	beginMinutes := string(beginTime[len(beginTime)-2]) // Grab last two digits of the time (minutes)
	beginTimeInt, _ := strconv.Atoi(beginTime)
	course.Assignment.HoursWeek = 3

	if isMTh {
		course.Assignment.EndTime = strconv.Itoa(beginTimeInt + 120)
		course.Assignment.Monday = true
		course.Assignment.Thursday = true
	} else {
		if beginMinutes == "00" {
			course.Assignment.EndTime = strconv.Itoa(beginTimeInt + 50)
		} else {
			course.Assignment.EndTime = strconv.Itoa(beginTimeInt + 90)
		}
		course.Assignment.Tuesday = true
		course.Assignment.Wednesday = true
		course.Assignment.Friday = true
	}

	if len(course.Assignment.EndTime) == 3 {
		course.Assignment.EndTime = "0" + course.Assignment.EndTime
	}

	return course
}

func setCourseDates(course structs.Course, term string) structs.Course {
	year := time.Now().Year()

	if term == "Fall" {
		course.Assignment.StartDate = "Sep 01, " + strconv.Itoa(year)
		course.Assignment.EndDate = "Dec 01, " + strconv.Itoa(year)
	} else if term == "Spring" {
		course.Assignment.StartDate = "Jan 01, " + strconv.Itoa(year+1)
		course.Assignment.EndDate = "Apr 01, " + strconv.Itoa(year+1)
	} else if term == "Summer" {
		course.Assignment.StartDate = "May 01, " + strconv.Itoa(year+1)
		course.Assignment.EndDate = "Aug 01, " + strconv.Itoa(year+1)
	} else {
		course.Assignment.StartDate = ""
		course.Assignment.EndDate = ""
	}

	return course
}
 
func ChangeRandomCourseTime(courses []structs.Course) []structs.Course {

    var randInt int

	timeslotMaps := CreateEmptyStreamType()

	AddCoursesToStreamMaps(courses, timeslotMaps)

	//choose course with large time gap for mutation
	for i := range courses {
		for j := range courses {
			if courses[i].StreamSequence == courses[j].StreamSequence {
				if courses[i].Assignment.Monday && courses[j].Assignment.Monday || courses[i].Assignment.Tuesday && courses[j].Assignment.Tuesday && courses[i].Prof.DisplayName != "TBD"{
					t1, err := strconv.Atoi(courses[i].Assignment.BeginTime)
					if err != nil {
						panic(err)
					}
					t2, err := strconv.Atoi(courses[j].Assignment.EndTime)
					if err != nil {
						panic(err)
					}
					if math.Copysign(float64(t1 - t2), 1) > 600 {
						randInt = j
						break
					}
				}
			}
		}
	}

	//add check for hard fixed courses?

	courses[randInt].Assignment.BeginTime = ""
	courses[randInt].Assignment.EndTime = ""

	courses[randInt].Assignment.Monday = false
	courses[randInt].Assignment.Tuesday = false
	courses[randInt].Assignment.Wednesday = false
	courses[randInt].Assignment.Thursday = false
	courses[randInt].Assignment.Friday = false


	if courses[randInt].StreamSequence == "1A" {
		courses[randInt], timeslotMaps.S1A, _ = addMultipleTimeslots(courses[randInt], timeslotMaps.S1A)
	} else if courses[randInt].StreamSequence == "1B" {
		courses[randInt], timeslotMaps.S1B, _ = addMultipleTimeslots(courses[randInt], timeslotMaps.S1B)
	} else if courses[randInt].StreamSequence == "2A" {
		courses[randInt], timeslotMaps.S2A, _ = addMultipleTimeslots(courses[randInt], timeslotMaps.S2A)
	} else if courses[randInt].StreamSequence == "2B" {
		courses[randInt], timeslotMaps.S2B, _ = addMultipleTimeslots(courses[randInt], timeslotMaps.S2B)
	} else if courses[randInt].StreamSequence == "3A" {
		courses[randInt], timeslotMaps.S3A, _ = addMultipleTimeslots(courses[randInt], timeslotMaps.S3A)
	} else if courses[randInt].StreamSequence == "3B" {
		courses[randInt], timeslotMaps.S3B, _ = addMultipleTimeslots(courses[randInt], timeslotMaps.S3B)
	} else if courses[randInt].StreamSequence == "4A" {
		courses[randInt], timeslotMaps.S4A, _ = addMultipleTimeslots(courses[randInt], timeslotMaps.S4A)
	} else if courses[randInt].StreamSequence == "4B" {
		courses[randInt], timeslotMaps.S4B, _ = addMultipleTimeslots(courses[randInt], timeslotMaps.S4B)
	}

	return courses
}

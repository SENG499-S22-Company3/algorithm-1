package scheduling

import "algorithm-1/structs"

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
			timeslotMaps.S1A = AddTimeslots(course, timeslotMaps.S1A)
		} else if course.StreamSequence == "1B" {
			timeslotMaps.S1B = AddTimeslots(course, timeslotMaps.S1B)
		} else if course.StreamSequence == "2A" {
			timeslotMaps.S2A = AddTimeslots(course, timeslotMaps.S2A)
		} else if course.StreamSequence == "2B" {
			timeslotMaps.S2B = AddTimeslots(course, timeslotMaps.S2B)
		} else if course.StreamSequence == "3A" {
			timeslotMaps.S3A = AddTimeslots(course, timeslotMaps.S3A)
		} else if course.StreamSequence == "3B" {
			timeslotMaps.S3B = AddTimeslots(course, timeslotMaps.S3B)
		} else if course.StreamSequence == "4A" {
			timeslotMaps.S4A = AddTimeslots(course, timeslotMaps.S4A)
		} else if course.StreamSequence == "4B" {
			timeslotMaps.S4B = AddTimeslots(course, timeslotMaps.S4B)
		}
	}

	return timeslotMaps
}

func AddTimeslots(course structs.Course, timeslots structs.Timeslots) structs.Timeslots {

	if course.Assignment.BeginTime != "" {
		if course.Assignment.Monday {
			timeslots.Monday[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
		}

		if course.Assignment.Tuesday {
			timeslots.Tuesday[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
		}
		if course.Assignment.Wednesday {
			timeslots.Wednesday[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
		}

		if course.Assignment.Thursday {
			timeslots.Thursday[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
		}

		if course.Assignment.Friday {
			timeslots.Friday[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
		}
	} else {
		// TO DO Handle non-historic courses
	}

	return timeslots
}

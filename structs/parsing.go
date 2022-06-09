package structs

import (
	"encoding/json"
	"regexp"
)

func ParseHistorical(jsonData []byte) Schedule {
	var parsedSchedule Schedule

	err := json.Unmarshal(jsonData, &parsedSchedule)

	if err != nil {
		panic(err)
	}

	return parsedSchedule
}

func Schedule2019(schedule Schedule) Schedule {
	var schedule2019 Schedule
	
	for _, x := range schedule.FallCourses {
		match, _ := regexp.MatchString("Sep .., 2018", x.Assignment.StartDate)
		if match {
			schedule2019.FallCourses = append(schedule2019.FallCourses, x)
		}
	}
	for _, x := range schedule.SpringCourses {
		match, _ := regexp.MatchString("Jan .., 2019", x.Assignment.StartDate)
		if match {
			schedule2019.SpringCourses = append(schedule2019.SpringCourses, x)
		}
	}
	for _, x := range schedule.SummerCourses {
		match, _ := regexp.MatchString("May .., 2019", x.Assignment.StartDate)
		if match {
			schedule2019.SummerCourses = append(schedule2019.SummerCourses, x)
		}
	}

	return schedule2019
}

func ParseCourses(jsonData []byte) []Course {
	var courses []Course
	err := json.Unmarshal(jsonData, &courses)
	if err != nil {
		panic(err)
	}

	return courses
}

func StructToJSON(schedule Schedule) []byte {

	jsonData, err := json.Marshal(schedule)

	if err != nil {
		panic(err)
	}

	return jsonData
}

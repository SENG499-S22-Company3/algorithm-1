package structs

import (
	"encoding/json"
	"regexp"
)

func ParseHistorical(jsonData []byte) (Schedule, error) {
	var parsedSchedule Schedule

	err := json.Unmarshal(jsonData, &parsedSchedule)

	if err != nil {
		return parsedSchedule, err
	}

	return parsedSchedule, nil
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

func ParseCourses(jsonData []byte) ([]Course, error) {
	var courses []Course
	err := json.Unmarshal(jsonData, &courses)
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func ParseProfPreferences(jsonData []byte) ([]Professor, error) {
	var profs []Professor

	err := json.Unmarshal(jsonData, &profs)

	if err != nil {
		return profs, err
	}

	return profs, nil
}

func StructToJSON(schedule Schedule) ([]byte, error) {

	jsonData, err := json.Marshal(schedule)

	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

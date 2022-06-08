package structs

import (
	"encoding/json"
)

func ParseHistorical(jsonData []byte) Schedule {
	var parsedSchedule Schedule

	err := json.Unmarshal(jsonData, &parsedSchedule)

	if err != nil {
		panic(err)
	}

	return parsedSchedule
}


func ParseCourses(jsonData []byte) []Course {
	var courses []Course
	err := json.Unmarshal(jsonData, &courses)
	if err != nil {
		panic(err)
	}

	return courses
}

func ParseProfPreferences(jsonData []byte) []Professor {
	var profs []Professor

	err := json.Unmarshal(jsonData, &profs)

	if err != nil {
		panic(err)
	}

	return profs
}


func StructToJSON(schedule Schedule) []byte {

	jsonData, err := json.Marshal(schedule)

	if err != nil {
		panic(err)
	}

	return jsonData
}


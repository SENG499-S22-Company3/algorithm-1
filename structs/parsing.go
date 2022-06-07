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

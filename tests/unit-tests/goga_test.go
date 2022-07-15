package tests

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGenetic(t *testing.T) {
	// preparing test data
	jsonData, err := ioutil.ReadFile("../data/input-test.json")
	if err != nil {
		t.Error("Error when opening input-test.json file: ", err)
	}

	input, err := structs.ParseInput(jsonData)
	if err != nil {
		t.Error("Input parsing failed with error: ", err.Error())
	}

	const term = "Fall"

	initSchedule := scheduling.Assignments(input.HardScheduled.FallCourses, input.CoursesToSchedule.FallCourses, input.Professors, "Fall")
	professors := append(input.Professors, structs.Professor{DisplayName: "TBD"})
	prefMap, _, teachingPrefMax := scheduling.MapPreferences(professors, term)

	startFit := int32(scheduling.GetFitness(initSchedule, prefMap, teachingPrefMax, term))
	fmt.Println("Starting Fitness: ", startFit)
	scheduling.PrettyPrintSemester(initSchedule)

	fmt.Println("starting ga test")
	var finalSchedule []structs.Course
	fit, i := -1, 0
	for int32(fit) <= startFit {

		timeslotMap, _ := scheduling.BaseTimeslotMaps(input.HardScheduled.FallCourses, term)
		requestedCourses, _, _ := scheduling.AddCoursesToStreamMaps(scheduling.Split(input.CoursesToSchedule.FallCourses), timeslotMap, "Fall")
		schedule := scheduling.AssignCourseProf(input.HardScheduled.FallCourses, requestedCourses, professors, term)
		scheduling.Optimize(schedule, professors, prefMap, teachingPrefMax, term)
		finalSchedule = append(schedule, input.HardScheduled.FallCourses...)
		fit = scheduling.GetFitness(finalSchedule, prefMap, teachingPrefMax, term)

		// timeout so GA doesn't take for so long
		if i > 60 {
			fmt.Println("BREAK")
			finalSchedule = initSchedule
			fit = scheduling.GetFitness(finalSchedule, prefMap, teachingPrefMax, term)
			break
		}
		i++
	}

	fmt.Println("ending ga test")
	scheduling.PrettyPrintSemester(finalSchedule)
	fmt.Println("Startfit: ", startFit)
	fmt.Println("Achieved Fitness: ", fit)
	fmt.Println("Max Fitness: ", (8*len(finalSchedule) + 32))
}

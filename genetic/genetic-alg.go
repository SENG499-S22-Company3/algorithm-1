package genetic

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"encoding/json"

	"github.com/MaxHalford/eaopt"

	"fmt"
)

var hardScheduled []structs.Course
var coursesToSchedule []structs.Course
var professors []structs.Professor
var term string

func RunGeneticAlg(inputHardScheduled []structs.Course, inputCourses []structs.Course, inputProfessors []structs.Professor, inputTerm string) ([]structs.Course, error) {

	hardScheduled = inputHardScheduled
	coursesToSchedule = inputCourses
	professors = inputProfessors
	term = inputTerm

	var ga, err = eaopt.NewDefaultGAConfig().NewGA()

	if err != nil {
		return nil, err
	}
	ga.NGenerations = 100
	ga.HofSize = 5
	ga.PopSize = 100
	ga.ParallelEval = true

	// Append the initial GA status to the progress file
	var bytes, _ = json.Marshal(ga)
	fmt.Println(string(bytes) + "\n")

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

	// Run the GA
	ga.Minimize(MakeSemester)

	fmt.Printf("Best GA with fitness score %f:\n", ga.HallOfFame[0].Fitness)

	jsonData, err := json.Marshal(ga.HallOfFame[0].Genome)

	if err != nil {
		return nil, err
	} else {
		goodSchedule, err := structs.ParseCourses(jsonData)
		if err != nil {
			return nil, err
		} else {
			_, err = scheduling.BaseTimeslotMaps(goodSchedule, term)
			if err != nil {
				return nil, err
			} else {
				err = scheduling.ScheduleConstraintsCheck(term, goodSchedule, professors)
				if err != nil {
					return nil, err
				} else {
					return goodSchedule, nil
				}
			}
		}
	}
}

func getInput() ([]structs.Course, []structs.Course, []structs.Professor, string) {
	return hardScheduled, coursesToSchedule, professors, term
}

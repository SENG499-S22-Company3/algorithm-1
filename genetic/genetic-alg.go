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
	ga.NGenerations = 20
	ga.HofSize = 5
	ga.PopSize = 500
	ga.ParallelEval = true

	// Append the initial GA status to the progress file
	var bytes, _ = json.Marshal(ga)
	fmt.Println(string(bytes) + "\n")

	// Add a custom print function to track progress
	/*
		ga.Callback = func(ga *eaopt.GA) {
			fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
		}
	*/

	// Run the GA
	ga.Minimize(MakeSemester)

	rerun := false
	hofSize := int(ga.HofSize)

	//runs the GA once more if no good schedule was found in Hall of fame
	for i := 0; i < hofSize; i++ {

		jsonData, err := json.Marshal(ga.HallOfFame[i].Genome)

		if err == nil {
			goodSchedule, err := structs.ParseCourses(jsonData)
			if err == nil {
				_, err = scheduling.BaseTimeslotMaps(goodSchedule, term)
				if err == nil {
					err = scheduling.ScheduleConstraintsCheck(term, goodSchedule, professors)
					if err == nil {
						return goodSchedule, nil
					}
				}
			}
		}

		if !rerun && i == hofSize-1 {
			ga.PopSize *= 2
			ga.NGenerations *= 2
			ga.Minimize(MakeSemester)
			rerun = true
			i = 0
		}

	}

	return nil, err

}

func getInput() ([]structs.Course, []structs.Course, []structs.Professor, string) {
	return hardScheduled, coursesToSchedule, professors, term
}

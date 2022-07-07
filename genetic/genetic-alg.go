package genetic

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"encoding/json"

	"github.com/MaxHalford/eaopt"

	"fmt"
)

func RunGeneticAlg() {

	var ga, err = eaopt.NewDefaultGAConfig().NewGA()

	if err != nil {
		fmt.Println(err)
		return
	}
	ga.NGenerations = 50
	ga.HofSize = 5
	ga.PopSize = 100

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
		fmt.Println(err)
		return
	} else {
		goodSchedule, err := structs.ParseCourses(jsonData)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			for _, course := range goodSchedule {
				fmt.Printf("%+v \n\n", course)
			}
			_, err = scheduling.BaseTimeslotMaps(goodSchedule)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Schedule Generated passes timeslot checks")
			}
		}
	}

	fmt.Println()
	fmt.Println()

}

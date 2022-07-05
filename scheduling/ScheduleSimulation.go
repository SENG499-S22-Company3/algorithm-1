package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"math/rand"

	ga "github.com/tomcraven/goga"
)

type ScheduleSimulation struct {
	simulationCount     int
	NumberOfSimulations int
	PopulationSize      int
	BaseSchedule        []structs.Course
	ProfList            []structs.Professor
}

// This function converts bits to a schedule
func NewSchedule(genome ga.Genome, courses []structs.Course, profs []structs.Professor) []structs.Course {
	bitset := genome.GetBits()

	//numberCourses = len(courses)
	//numberProfs = len(profs)

	times := []string{"0830", "1000", "1130", "1300"}

	bits := bitset.GetAll()
	for i, j := 0, 0; i < len(bits); i, j = i+4, j+1 {
		var assignment []int
		if i > len(bits)-4 {
			assignment = bits[i:]
		} else {
			assignment = bits[i : i+4]
		}

		// decoding 4 bits into prof and timesot
		profIndex := assignment[0]*2 + assignment[1]
		timeIndex := assignment[2]*2 + assignment[3]

		courses[j].Prof = profs[profIndex]
		time := times[timeIndex]
		courses[j].Assignment = structs.Assignment{
			Monday:    true,
			BeginTime: time,
		}

	}

	return courses
}

// Go initializes a random roster
func (s ScheduleSimulation) Go() ga.Bitset {
	//size := s.NumberOfCourses * s.NumberOfProfs * 3
	size := 16 // this is long enough to fit 4 4bit assignments (2 bits for prof and 2 bits fo time)
	bitset := ga.Bitset{}
	bitset.Create(size)
	for i := 0; i < size; i++ {
		bitset.Set(i, rand.Intn(2))
	}
	return bitset
}

func (sim *ScheduleSimulation) OnBeginSimulation() {
	sim.simulationCount++
	if sim.NumberOfSimulations < 1 {
		panic("NumberOfSimulations must be greater than 0")
	}
}

// Simulate assigns a fitness value to the given genome
func (sim *ScheduleSimulation) Simulate(genome ga.Genome) {
	schedule := NewSchedule(genome, sim.BaseSchedule, sim.ProfList)
	fitness := GetFitness(schedule)
	(genome).SetFitness(fitness)
}

func GetFitness(s []structs.Course) int {
	// hard coded scoring for example
	score := 0
	// brute force that the times are different
	if s[0].Assignment.BeginTime != s[1].Assignment.BeginTime {
		score += 10
	}
	if s[0].Assignment.BeginTime != s[2].Assignment.BeginTime {
		score += 10
	}
	if s[0].Assignment.BeginTime != s[3].Assignment.BeginTime {
		score += 10
	}
	if s[1].Assignment.BeginTime != s[2].Assignment.BeginTime {
		score += 10
	}
	if s[1].Assignment.BeginTime != s[3].Assignment.BeginTime {
		score += 10
	}
	if s[2].Assignment.BeginTime != s[3].Assignment.BeginTime {
		score += 10
	}

	// brute force the profs are all different
	if s[0].Prof.DisplayName != s[1].Prof.DisplayName {
		score += 10
	}
	if s[0].Prof.DisplayName != s[2].Prof.DisplayName {
		score += 10
	}
	if s[0].Prof.DisplayName != s[3].Prof.DisplayName {
		score += 10
	}
	if s[1].Prof.DisplayName != s[2].Prof.DisplayName {
		score += 10
	}
	if s[1].Prof.DisplayName != s[3].Prof.DisplayName {
		score += 10
	}
	if s[2].Prof.DisplayName != s[3].Prof.DisplayName {
		score += 10
	}

	// adding preference to score
	score += int(s[3].Prof.Preferences[0].PreferenceNum)

	return score
}

// OnElite prints the current elite on every simulation iteration
func (sim *ScheduleSimulation) OnElite(genome ga.Genome) {
	schedule := NewSchedule(genome, sim.BaseSchedule, sim.ProfList)

	fmt.Println("***********************")
	fmt.Printf("** [%d] simulation **\n", sim.simulationCount)
	fmt.Println("solution: ")
	//fmt.Println(schedule)
	prettyPrintSemester(schedule)
	//fmt.Print("bits: ")
	//fmt.Println(genome.GetBits().GetAll())
	fmt.Print("fitness: ")
	fmt.Println(GetFitness(schedule))
	fmt.Println("***********************")
}

func prettyPrintSemester(s []structs.Course) {
	for _, c := range s {
		fmt.Print(c.Subject)
		fmt.Print(c.CourseNumber)
		fmt.Print(", ")
		fmt.Print(c.Prof.DisplayName)
		fmt.Print(", ")
		fmt.Print(c.Assignment.BeginTime)
		fmt.Println()
		fmt.Println("---------------------------------")
	}
}

// ExitFunc defines when to stop the simulation
func (sim *ScheduleSimulation) ExitFunc(genome ga.Genome) bool {
	return sim.simulationCount >= sim.NumberOfSimulations
}

func (sim *ScheduleSimulation) OnEndSimulation() {}

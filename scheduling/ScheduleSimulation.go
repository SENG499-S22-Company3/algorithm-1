package scheduling

import (
	"fmt"
	"math"
	"math/rand"

	ga "github.com/tomcraven/goga"
)

type ScheduleSimulation struct {
	simulationCount     int
	NumberOfSimulations int
	NumberOfProfs       int
	NumberOfCourses     int
	PopulationSize      int
}

// this probably shouldn't really go here?
func NewSchedule(genome ga.Genome, employees int, days int) int {
	// for each employee go over the bitset and extract the current shift assigment for the whole period of time
	bits := genome.GetBits()
	return bitsToStruct(bits)
}

func setUpBitset() {

}

// will need to actually convert bits to a schedule
func bitsToStruct(bitset *ga.Bitset) int {
	/*
		c1 := structs.Course{
			CourseNumber:   "310",
			Subject:        "SENG",
			SequenceNumber: "A01",
			StreamSequence: "2A",
			CourseTitle:    "Human Computer Interaction",
		}
		c2 := structs.Course{
			CourseNumber:   "370",
			Subject:        "CSC",
			SequenceNumber: "A01",
			StreamSequence: "2A",
			CourseTitle:    "Database Systems",
		}
		c3 := structs.Course{
			CourseNumber:   "361",
			Subject:        "CSC",
			SequenceNumber: "A01",
			StreamSequence: "2A",
			CourseTitle:    "Computer Communications and Networks",
		}
		p1 := structs.Professor{
			DisplayName: "A",
		}
		p2 := structs.Professor{
			DisplayName: "B",
		}
		p3 := structs.Professor{
			DisplayName: "C",
		}
		slots := structs.Timeslots{
			Monday: map[string]string{
				"0830": "",
				"1000": "",
				"1130": "",
			},
		}
		courses := []structs.Course{c1, c2, c3}
		profs := []structs.Professor{p1, p2, p3}
	*/

	bits := bitset.GetAll()
	sum := 0
	scale := 512
	for _, bit := range bits {
		sum += bit * scale
		scale /= 2
	}
	return sum
}

// Go initializes a random roster
func (s ScheduleSimulation) Go() ga.Bitset {
	//size := s.NumberOfCourses * s.NumberOfProfs * 3
	size := 10
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
	schedule := NewSchedule(genome, sim.NumberOfProfs, sim.NumberOfCourses)
	fitness := GetFitness(schedule)
	(genome).SetFitness(fitness)
}

func GetFitness(s int) int {
	//return modclosestto(s)
	return matchNumber(s)
}

// example fitness for numbers
func matchNumber(s int) int {
	score := 100.0
	target := 3
	score -= math.Abs((float64((s - target))))
	return int(score)
}

// example fitness for numbers
func modclosestto(s int) int {
	score := 0.0
	if s%7 == 0 {
		score = 100
	}
	score -= math.Abs((float64((s - 500))))
	return int(score)
}

// OnElite prints the current elite on every simulation iteration
func (r *ScheduleSimulation) OnElite(genome ga.Genome) {
	schedule := NewSchedule(genome, r.NumberOfProfs, r.NumberOfCourses)
	fmt.Printf("** [%d] simulation **\n", r.simulationCount)
	fmt.Println(schedule)
	fmt.Println(genome.GetBits().GetAll())
	fmt.Println(GetFitness(schedule))
}

// ExitFunc defines when to stop the simulation
func (sim *ScheduleSimulation) ExitFunc(genome ga.Genome) bool {
	return sim.simulationCount >= sim.NumberOfSimulations
}

func (sim *ScheduleSimulation) OnEndSimulation() {}

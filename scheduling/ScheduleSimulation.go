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
	NumberOfProfs       int
	NumberOfCourses     int
	PopulationSize      int
}

// this probably shouldn't really go here?
func NewSchedule(genome ga.Genome, employees int, days int) []structs.Course {
	// for each employee go over the bitset and extract the current shift assigment for the whole period of time
	bits := genome.GetBits()
	return bitsToStruct(bits)
}

func setUpBitset() {

}

// will need to actually convert bits to a schedule
func bitsToStruct(bitset *ga.Bitset) []structs.Course {
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
	c4 := structs.Course{
		CourseNumber:   "320",
		Subject:        "CSC",
		SequenceNumber: "A01",
		StreamSequence: "2A",
		CourseTitle:    "Fundamentals of Computer Science",
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
	p4 := structs.Professor{
		DisplayName: "D",
	}

	times := []string{"0830", "1000", "1130", "1300"}
	courses := []structs.Course{c1, c2, c3, c4}
	profs := []structs.Professor{p1, p2, p3, p4}

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
	schedule := NewSchedule(genome, sim.NumberOfProfs, sim.NumberOfCourses)
	fitness := GetFitness(schedule)
	(genome).SetFitness(fitness)
}

func GetFitness(s []structs.Course) int {
	score := 0
	if s[0].CourseNumber == "310" {
		score += 10
	}
	if s[1].CourseNumber == "370" {
		score += 10
	}
	if s[2].CourseNumber == "361" {
		score += 10
	}

	return score
}

// OnElite prints the current elite on every simulation iteration
func (r *ScheduleSimulation) OnElite(genome ga.Genome) {
	schedule := NewSchedule(genome, r.NumberOfProfs, r.NumberOfCourses)
	fmt.Printf("** [%d] simulation **\n", r.simulationCount)
	fmt.Print("solution: ")
	fmt.Println(schedule)
	//fmt.Print("bits: ")
	//fmt.Println(genome.GetBits().GetAll())
	fmt.Print("fitness: ")
	fmt.Println(GetFitness(schedule))
}

// ExitFunc defines when to stop the simulation
func (sim *ScheduleSimulation) ExitFunc(genome ga.Genome) bool {
	return sim.simulationCount >= sim.NumberOfSimulations
}

func (sim *ScheduleSimulation) OnEndSimulation() {}

package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"math"
	"math/rand"

	ga "github.com/tomcraven/goga"
)

type ScheduleSimulation struct {
	simulationCount     int
	NumberOfSimulations int
	PopulationSize      int
	BaseSemester        []structs.Course
	NumberOfCourses     int
	ProfList            []structs.Professor
	NumberOfProfs       int
	SectionBitWidth     int
}

// timeslots: 1 bit for day | 4 bits for time
const timeslotBitWidth = 5

// This function converts bits to a schedule
func NewSchedule(genome ga.Genome, sim ScheduleSimulation) []structs.Course {
	courses := sim.BaseSemester
	profs := sim.ProfList
	times := []string{"0830", "1000", "1130", "1300", "1430", "1600", "1730",
		"0930", "1030", "1130", "1230", "1330", "1430", "1530", "1630", "1730"}

	bits := genome.GetBits().GetAll()
	for i, j := 0, 0; i < len(bits); i, j = i+sim.SectionBitWidth, j+1 {
		assignment := bits[i : i+sim.SectionBitWidth]

		// decoding section of bits into timeslot and prof indexes
		// first bit determines day A or day B
		dayIndex := assignment[0]
		// next few bits are the timeslot
		timeIndex := bitsToNumber(assignment[1:timeslotBitWidth])
		// rest of the bits are the prof
		profIndex := bitsToNumber(assignment[timeslotBitWidth:])

		// need to be careful of invalid indexes
		if profIndex < sim.NumberOfProfs {
			courses[j].Prof = profs[profIndex]
		}

		time := times[timeIndex]

		courses[j].Assignment = structs.Assignment{
			Monday:    dayIndex == 1,
			Tuesday:   dayIndex == 0,
			Wednesday: dayIndex == 0,
			Thursday:  dayIndex == 1,
			Friday:    dayIndex == 0,
			BeginTime: time,
		}
	}

	return courses
}

// takes array of bits and turns it into an int
// ex {1,1,1,1} becomes 15 since 1111 in binary is 15 in decimal
func bitsToNumber(bits []int) int {
	length := len(bits)
	scale := int(math.Pow(2, float64(length-1)))
	sum := 0
	for _, b := range bits {
		sum += scale * b
		scale /= 2
	}
	return sum
}

func scheduleToBitset(semester []structs.Course, profs []structs.Professor) []int {
	var times = map[string]int{
		"0830" : 0, "1000" : 1, "1130" : 2, "1300" : 3, "1430" : 4, "1600" : 5, "1730" : 6,
		"0930" : 7, "1030" : 8, "1230" : 10, "1330" : 11, "1530" : 13, "1630" : 14,
	}

	var bits []int
	for _, course := range(semester) {
		var day int
		if course.Assignment.Monday {
			day = 1
		} else {
			day = 0
		}

		time := numberToBits(times[course.Assignment.BeginTime])

		prof := numberToBits(getProfIndex(course.Prof.DisplayName, profs))

		bits = append(bits, day)
		bits = append(bits, time...)
		bits = append(bits, prof...)
	}

	return bits
}

func numberToBits(num int) []int {
	var bitset []int
	binary := fmt.Sprintf("%b", num)
	for i, bit := range(binary) {
		bitset[i] = int(bit) - '0'
	}
	return bitset
}

func getProfIndex(name string, profs []structs.Professor) int {
	for i, p := range(profs) {
		if p.DisplayName == name {
			return i
		}
	}
	return -1
}

// Initializes a random schedule to start with
func (sim ScheduleSimulation) Go() ga.Bitset {
	size := sim.NumberOfCourses * sim.SectionBitWidth
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
	schedule := NewSchedule(genome, *sim)
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
	schedule := NewSchedule(genome, *sim)

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
		fmt.Print(c.CourseNumber, "\t")
		fmt.Print(c.Assignment.BeginTime, "\t")
		fmt.Print(c.Assignment.Monday, " ")
		fmt.Print(c.Assignment.Tuesday, " ")
		fmt.Print(c.Assignment.Wednesday, " ")
		fmt.Print(c.Assignment.Thursday, " ")
		fmt.Print(c.Assignment.Friday, "\t")
		fmt.Print(c.Prof.DisplayName)
		fmt.Println()
	}
}

// ExitFunc defines when to stop the simulation
func (sim *ScheduleSimulation) ExitFunc(genome ga.Genome) bool {
	return sim.simulationCount >= sim.NumberOfSimulations
}

func (sim *ScheduleSimulation) OnEndSimulation() {}

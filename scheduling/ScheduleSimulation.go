package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"math"
	"sort"

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
	PreferenceMap       map[string]map[string]int
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

func scheduleToBitset(sim ScheduleSimulation) []int {
	semester := sim.BaseSemester
	profs := sim.ProfList
	var times = map[string]int{
		"0830": 0, "1000": 1, "1130": 2, "1300": 3, "1430": 4, "1600": 5, "1730": 6,
		"0930": 7, "1030": 8, "1230": 10, "1330": 11, "1530": 13, "1630": 14,
	}

	var bits []int
	for _, course := range semester {
		var day int
		if course.Assignment.Monday {
			day = 1
		} else {
			day = 0
		}

		time := numberToBits(times[course.Assignment.BeginTime], 4)

		prof := numberToBits(getProfIndex(course.Prof.DisplayName, profs), sim.SectionBitWidth-timeslotBitWidth)

		bits = append(bits, day)
		bits = append(bits, time...)
		bits = append(bits, prof...)
	}

	return bits
}

func numberToBits(num int, length int) []int {
	var bitset []int
	binary := fmt.Sprintf("%b", num) // turns an int into binary string
	for i := 0; i < len(binary); i++ {
		bitset = append(bitset, int(binary[i])-'0') // turning the string into []int
	}
	for len(bitset) != length {
		bitset = append([]int{0}, bitset...) // prepending 0's to fit desired length
	}
	return bitset
}

// returns the index in the given slice where the prof with the given name is located
func getProfIndex(name string, profs []structs.Professor) int {
	for i, p := range profs {
		if p.DisplayName == name {
			return i
		}
	}
	return -1 // returns -1 if the prof is not found
}

// Initializes the bitset with an initial schedule
func (sim ScheduleSimulation) Go() ga.Bitset {
	size := sim.NumberOfCourses * sim.SectionBitWidth
	bitset := ga.Bitset{}
	bitset.Create(size)
	//if sim.c
	bits := scheduleToBitset(sim)
	for i := 0; i < size; i++ {
		bitset.Set(i, bits[i])
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
	fitness := GetFitness(schedule, sim)
	(genome).SetFitness(fitness)
}

func GetFitness(s []structs.Course, sim *ScheduleSimulation) int {
	// hard coded scoring for example
	score := 0

	var teachingMap = map[string]string{}
	timeslotMaps, _ := BaseTimeslotMaps(s)

	for _, c := range s {

		// professors checks
		prof := c.Prof.DisplayName
		var d string
		if c.Assignment.Monday {
			d = "MTh" + c.Assignment.BeginTime
		} else {
			d = "TWF" + c.Assignment.BeginTime
		}

		if _, timeConflict := teachingMap[prof+d]; timeConflict {
			return 0
		}
		teachingMap[prof+d] = c.CourseTitle
		score += int(sim.PreferenceMap[prof][(c.Subject + c.CourseNumber)])

		// timeslots checks
		if c.StreamSequence == "1A" {
			score, timeslotMaps.S1A = checkStream(c, timeslotMaps.S1A, score)
		} else if c.StreamSequence == "1B" {
			score, timeslotMaps.S1B = checkStream(c, timeslotMaps.S1B, score)
		} else if c.StreamSequence == "2A" {
			score, timeslotMaps.S2A = checkStream(c, timeslotMaps.S2A, score)
		} else if c.StreamSequence == "2B" {
			score, timeslotMaps.S2B = checkStream(c, timeslotMaps.S2B, score)
		} else if c.StreamSequence == "3A" {
			score, timeslotMaps.S3A = checkStream(c, timeslotMaps.S3A, score)
		} else if c.StreamSequence == "3B" {
			score, timeslotMaps.S3B = checkStream(c, timeslotMaps.S3B, score)
		} else if c.StreamSequence == "4A" {
			score, timeslotMaps.S4A = checkStream(c, timeslotMaps.S4A, score)
		} else if c.StreamSequence == "4B" {
			score, timeslotMaps.S4B = checkStream(c, timeslotMaps.S4B, score)
		}

		if score == 0 {
			return 0
		}
	}

	return score
}

func checkStream(course structs.Course, timeslots structs.Timeslots, score int) (int, structs.Timeslots) {
	if course.Assignment.Monday {
		// fmt.Println("Monday:",  timeslots.Monday)
		timeslots.Monday, score = checkConflict(course, timeslots.Monday, score)
	}
	if course.Assignment.Tuesday {
		// fmt.Println("Tuesday:", timeslots.Tuesday)
		timeslots.Tuesday, score = checkConflict(course, timeslots.Tuesday, score)
	}
	if course.Assignment.Wednesday {
		// fmt.Println("Wednesday:", timeslots.Wednesday)
		timeslots.Wednesday, score = checkConflict(course, timeslots.Wednesday, score)
	}
	if course.Assignment.Thursday {
		// fmt.Println("Thursday:", timeslots.Thursday)
		timeslots.Thursday, score = checkConflict(course, timeslots.Thursday, score)
	}
	if course.Assignment.Friday {
		// fmt.Println("Friday:", timeslots.Friday)
		timeslots.Friday, score = checkConflict(course, timeslots.Friday, score)
	}
	return score, timeslots
}

func checkConflict(course structs.Course, day map[string]string, score int) (map[string]string, int) {

	if _, isValid := day[course.Assignment.BeginTime]; !isValid { // Check if map key exists
		fmt.Println("Error not valid")
		score = 0
	} else if scheduledCourse := day[course.Assignment.BeginTime]; scheduledCourse != "" { // Check if there is already a course there
		fmt.Println("Error already scheduled")
		score = 0
	} else {
		day[course.Assignment.BeginTime] = course.Subject + course.CourseNumber
		score = score + 1
	}

	return day, score
}

// OnElite prints the current elite on every simulation iteration
func (sim *ScheduleSimulation) OnElite(genome ga.Genome) {
	schedule := NewSchedule(genome, *sim)

	fmt.Println("***********************")
	fmt.Printf("** [%d] simulation **\n", sim.simulationCount)
	fmt.Println("solution: ")
	prettyPrintSemester(schedule)
	fmt.Print("fitness: ")
	fmt.Println(GetFitness(schedule, sim))
	fmt.Println("***********************")
}

func prettyPrintSemester(s []structs.Course) {

	sort.SliceStable(s, func(i, j int) bool {
		return s[i].StreamSequence < s[j].StreamSequence
	})

	for i, c := range s {
		fmt.Print(i, "\t", c.Subject)
		fmt.Print(c.CourseNumber, "\t")
		fmt.Print(c.SequenceNumber, " ")
		fmt.Print(c.StreamSequence, "\t")
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

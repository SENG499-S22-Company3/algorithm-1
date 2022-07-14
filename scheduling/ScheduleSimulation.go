package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"math"
	"strconv"

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
	TeachingPrefMax		map[string]int
}

// timeslots: 4 bits for time
const timeslotBitWidth = 4

// This function converts bits to a schedule
func NewSchedule(genome ga.Genome, sim ScheduleSimulation) []structs.Course {
	courses := sim.BaseSemester
	profs := sim.ProfList
	begin_times := []string{"0830", "1000", "1130", "1300", "1430", "1600", "1730",
		"0830", "0930", "1030", "1130", "1230", "1330", "1430", "1530", "1630"}
	end_times := []string{"0950", "1120", "1250", "1420", "1550", "1720", "1820",
		"0920", "1020", "1120", "1220", "1320", "1420", "1520", "1620", "1720"}

	bits := genome.GetBits().GetAll()
	for i, j := 0, 0; i < len(bits); i, j = i+sim.SectionBitWidth, j+1 {
		assignment := bits[i : i+sim.SectionBitWidth]

		// decoding section of bits into timeslot and prof indexes
		// first few bits are the timeslot
		timeIndex := bitsToNumber(assignment[:timeslotBitWidth])
		// rest of the bits are the prof
		profIndex := bitsToNumber(assignment[timeslotBitWidth:])
		// determines day A or day B
		dayA := timeIndex < 6
		dayB := !dayA

		// fmt.Println("profIndex", profIndex)

		// need to be careful of invalid indexes
		if profIndex < sim.NumberOfProfs {
			courses[j].Prof = profs[profIndex]
		}

		courses[j].Assignment = structs.Assignment{
			Monday:    dayA,
			Tuesday:   dayB,
			Wednesday: dayB,
			Thursday:  dayA,
			Friday:    dayB,
			BeginTime: begin_times[timeIndex],
			EndTime:   end_times[timeIndex],
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
		"0830A": 0,
		"1000A": 1,
		"1130A": 2,
		"1300A": 3,
		"1430A": 4,
		"1600A": 5,
		"1730A": 6,
		"0830B": 7,
		"0930B": 8,
		"1030B": 9,
		"1130B": 10,
		"1230B": 11,
		"1330B": 12,
		"1430B": 13,
		"1530B": 14,
		"1630B": 15,
	}

	var bits []int
	for _, course := range semester {
		var day string
		if course.Assignment.Monday {
			day = "A"
		} else {
			day = "B"
		}

		time := numberToBits(times[course.Assignment.BeginTime+day], 4)

		prof := numberToBits(getProfIndex(course.Prof.DisplayName, profs), sim.SectionBitWidth-timeslotBitWidth)

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
	fitness := GetFitness(schedule, sim.PreferenceMap, sim.TeachingPrefMax)
	// fmt.Println("Fitnes: ", fitness)
	(genome).SetFitness(fitness)
}

type TimeMinMax struct {
	StartMin string
	EndMax   string
}

func GetFitness(s []structs.Course, prefMap map[string]map[string]int, TeachingPrefMax map[string]int) int {
	score := 0

	// timeslot checks
	if _, err := BaseTimeslotMaps(s); err != nil {
		return 0
	} else {
		score += len(s) * 2
	}

	// professors checks
	var teachingMap = map[string]string{}
	var teachingCount = map[string]int{}
	var timeMap = map[string]TimeMinMax{}
	for _, c := range s {

		var days string
		var beginTime string
		if c.Assignment.Monday {
			days = "MTh"
			beginTime = c.Assignment.BeginTime
		} else {
			days = "TWF"
			beginTime = c.Assignment.BeginTime
		}

		prof := c.Prof.DisplayName
		score += int(prefMap[prof][(c.Subject + c.CourseNumber)])
		if _, timeConflict := teachingMap[prof+days+beginTime]; timeConflict {
			return 0
		}

		if prof != "TBD" {
			teachingMap[prof+days+beginTime] = c.CourseTitle
			teachingCount[prof]++
			if teachingCount[prof] > TeachingPrefMax[prof] {
				return 0
			}
		}

		// track min and max for every stream sequence for both MTh and TWF
		if _, keyExists := timeMap[c.StreamSequence+days]; keyExists {
			if c.Assignment.BeginTime < timeMap[c.StreamSequence+days].StartMin {
				timeMap[c.StreamSequence+days] = TimeMinMax{
					StartMin: c.Assignment.BeginTime,
					EndMax:   timeMap[c.StreamSequence+days].EndMax,
				}
			}

			if c.Assignment.EndTime > timeMap[c.StreamSequence+days].EndMax {
				timeMap[c.StreamSequence+days] = TimeMinMax{
					StartMin: timeMap[c.StreamSequence+days].StartMin,
					EndMax:   c.Assignment.EndTime,
				}
			}
		} else {
			timeMap[c.StreamSequence+days] = TimeMinMax{
				StartMin: c.Assignment.BeginTime,
				EndMax:   c.Assignment.EndTime,
			}
		}

	}

	// add score to streams that
	for _, element := range timeMap {
		EndMax, _ := strconv.Atoi(element.EndMax)
		startMin, _ := strconv.Atoi(element.StartMin)

		if (EndMax - startMin) < 600 {
			score += 2
		}
	}

	return score

}

// OnElite prints the current elite on every simulation iteration
func (sim *ScheduleSimulation) OnElite(genome ga.Genome) {
	// schedule := NewSchedule(genome, *sim)

	// fmt.Println("***********************")
	// fmt.Printf("** [%d] simulation **\n", sim.simulationCount)
	// // // fmt.Println("solution: ")
	// //PrettyPrintSemester(schedule)
	// fmt.Print("fitness: ")
	// fmt.Println(GetFitness(schedule, sim.PreferenceMap))
	// fmt.Println("***********************")
}

func PrettyPrintSemester(s []structs.Course) {

	// sort.SliceStable(s, func(i, j int) bool {
	// 	return s[i].StreamSequence < s[j].StreamSequence
	// })

	for i, c := range s {
		fmt.Print(i, "\t", c.Subject)
		fmt.Print(c.CourseNumber, "\t")
		fmt.Print(c.SequenceNumber, " ")
		fmt.Print(c.StreamSequence, "\t")
		fmt.Print(c.Assignment.BeginTime, " to ")
		fmt.Print(c.Assignment.EndTime, "\t")
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

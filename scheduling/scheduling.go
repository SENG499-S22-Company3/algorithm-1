package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"runtime"
	"time"

	ga "github.com/tomcraven/goga"
)

func sliceContains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func BaseSemester(requestedCourses []structs.Course, hardSemester []structs.Course) []structs.Course {
	var result []structs.Course
	ignore := []string{"CHEM", "MATH", "PHYS", "STAT", "ECON"} // Not sure if this is comprehensive
	for _, h := range hardSemester {
		if sliceContains(h.Subject, ignore) {
			result = append(result, h)
		}
		for _, c := range requestedCourses {
			if c.Subject == h.Subject && c.CourseNumber == h.CourseNumber {
				h.StreamSequence = c.StreamSequence
				result = append(result, h)
			}
		}
	}
	return result
}

func BaseSchedule(requestedCourses structs.Schedule, hardSchedule structs.Schedule) structs.Schedule {
	// making channels to get return values from goroutines
	fall := make(chan []structs.Course)
	spring := make(chan []structs.Course)
	summer := make(chan []structs.Course)

	go func() {
		fall <- BaseSemester(requestedCourses.FallCourses, hardSchedule.FallCourses)
	}()
	go func() {
		spring <- BaseSemester(requestedCourses.SpringCourses, hardSchedule.SpringCourses)
	}()
	go func() {
		summer <- BaseSemester(requestedCourses.SummerCourses, hardSchedule.SummerCourses)
	}()

	return structs.Schedule{
		FallCourses:   <-fall,
		SpringCourses: <-spring,
		SummerCourses: <-summer,
	}
}

func Assignments(hardScheduledCourses []structs.Course, requestedCourses []structs.Course, professors []structs.Professor, term string) []structs.Course {

	timeslotMap, _ := BaseTimeslotMaps(hardScheduledCourses, term)
	requestedCourses, _, _ = AddCoursesToStreamMaps(Split(requestedCourses), timeslotMap, term)
	requestedCourses = AssignCourseProf(hardScheduledCourses, requestedCourses, professors, term)
	requestedCourses = append(requestedCourses, hardScheduledCourses...)

	return requestedCourses
}

func Optimize(schedule structs.Schedule, professors []structs.Professor) {
	// simulation for fall semester
	simulation := ScheduleSimulation{
		NumberOfSimulations: 100,
		PopulationSize:      20,
		BaseSchedule:        schedule.FallCourses,
		ProfList:            professors,
	}

	// mater defines how to combine genomes
	mater := ga.NewMater(
		[]ga.MaterFunctionProbability{
			{P: 1.0, F: ga.TwoPointCrossover},
			{P: 1.0, F: ga.Mutate},
			{P: 1.0, F: ga.UniformCrossover, UseElite: true},
		},
	)

	// selector defines how to select genomes from which the elite is being taken
	selector := ga.NewSelector(
		[]ga.SelectorFunctionProbability{
			{P: 1.0, F: ga.Roulette},
		},
	)

	algorithm := ga.NewGeneticAlgorithm()
	algorithm.Simulator = &simulation
	algorithm.EliteConsumer = &simulation
	algorithm.BitsetCreate = &simulation
	algorithm.Selector = selector
	algorithm.Mater = mater

	numThreads := 4
	runtime.GOMAXPROCS(numThreads)
	algorithm.Init(simulation.PopulationSize, numThreads)

	startTime := time.Now()
	algorithm.Simulate()
	fmt.Println(time.Since(startTime))

}

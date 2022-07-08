package genetic

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"math/rand"
	"strings"
	//"fmt"
	"strconv"
	"math"
	"github.com/MaxHalford/eaopt"
	"github.com/jinzhu/copier"
)

// A Semester is a slice of Courses.
type Semester []structs.Course

// At method from Slice
func (sem Semester) At(i int) interface{} {
	return sem[i]
}

// Set method from Slice
func (sem Semester) Set(i int, v interface{}) {
	//check for deep copy
	sem[i] = v.(structs.Course)
}

// Len method from Slice
func (sem Semester) Len() int {
	return len(sem)
}

// Swap method from Slice
func (sem Semester) Swap(i, j int) {
	//deep copy?
	sem[i], sem[j] = sem[j], sem[i]
}

// Slice method from Slice
func (sem Semester) Slice(a, b int) eaopt.Slice {
	//deep copy
	return sem[a:b]
}

// Split method from Slice
func (sem Semester) Split(k int) (eaopt.Slice, eaopt.Slice) {
	//deep copy
	return sem[:k], sem[k:]
}

// TODO Append method from Slice - should not be used?
func (sem Semester) Append(q eaopt.Slice) eaopt.Slice {
	return append(sem, q.(Semester)...)
	//return sem
}

// Replace method from Slice
func (sem Semester) Replace(q eaopt.Slice) {
	copy(sem, q.(Semester))
}

// Copy method from Slice
func (sem Semester) Copy() eaopt.Slice {
	clone := sem
	copier.Copy(&clone, &sem)
	return clone
}

// Clone a Semester.
func (sem Semester) Clone() eaopt.Genome {
	var clone = make(Semester, len(sem))
	//clone := sem
	//copy(clone, sem)
	copier.Copy(&clone, &sem)
	return clone

}

// Evaluate a Semester by summing the consecutive Euclidean distances.
func (sem Semester) Evaluate() (penalty float64, err error) {

	//var courses []structs.Course
	var profCourseCount int

	//evalute prof clashes
	for i := range sem {
		//courses = append(courses, sem[i])
		profCourseCount = 0
		for j := range sem {
			//evalute prof clashes
			if sem[i].Prof.DisplayName == sem[j].Prof.DisplayName && i != j && sem[i].Prof.DisplayName != "TBD"{
				if sem[i].Assignment.BeginTime == sem[j].Assignment.BeginTime {
					if sem[i].Assignment.Monday && sem[j].Assignment.Monday || sem[i].Assignment.Tuesday && sem[j].Assignment.Tuesday {
						penalty += 1000
					}
				}
				//penalize profs teaching > 3 courses
				profCourseCount += 1
				if profCourseCount > 3{
					penalty += 1000
				}
			}
			//evaluate same stream time differences
			if sem[i].StreamSequence == sem[j].StreamSequence {
				if sem[i].Assignment.Monday && sem[j].Assignment.Monday || sem[i].Assignment.Tuesday && sem[j].Assignment.Tuesday && sem[i].Prof.DisplayName != "TBD"{
					t1, err := strconv.Atoi(sem[i].Assignment.BeginTime)
					if err != nil {
						panic(err)
					}
					t2, err := strconv.Atoi(sem[j].Assignment.EndTime)
					if err != nil {
						panic(err)
					}
					if math.Copysign(float64(t1 - t2), 1) > 600 {
						penalty += 300
					}
				}
			}
		}
		//penalize low prof preference values
		for j := range sem[i].Prof.Preferences {
			if strings.Contains(sem[i].Prof.Preferences[j].CourseNum, sem[i].CourseNumber) && strings.Contains(sem[i].Prof.Preferences[j].CourseNum, sem[i].Subject) {
				penalty += 10 * (6 - float64(sem[i].Prof.Preferences[j].PreferenceNum))
			}
		}
	}

	// Check if timeslots violate hard requirements
	_, fail := scheduling.BaseTimeslotMaps(sem)
	if fail != nil {
		penalty += 1000
	}

	return
}

// Mutate a Semester by applying by permutation mutation and/or splice mutation.
func (sem Semester) Mutate(rng *rand.Rand) {

	if rng.Float64() < 0.45 {
		sem = scheduling.ChangeRandomCourseTime(sem)
		sem = scheduling.ChangeRandomCourseProf(sem)
	}
}

// Crossover a Semester with another Semester by using Partially Mixed Crossover (PMX).
func (sem Semester) Crossover(q eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNX(sem, q.(Semester), 3, rng)
}

// MakeSemester creates a random semester
func MakeSemester(rng *rand.Rand) eaopt.Genome {

	input := getInput()

	testStreamtype, _ := scheduling.BaseTimeslotMaps(input.HardScheduled.SpringCourses)
	input.CoursesToSchedule.SpringCourses, _, _ = scheduling.AddCoursesToStreamMaps(input.CoursesToSchedule.SpringCourses, testStreamtype)
	testScheduleCourse := scheduling.AssignCourseProf(input.CoursesToSchedule.SpringCourses, input.CoursesToSchedule.SpringCourses, input.Professors)
	testScheduleCourse = append(testScheduleCourse, input.HardScheduled.SpringCourses...)

	testSem := make(Semester, len(testScheduleCourse))

	copy(testSem, testScheduleCourse)

	return testSem
}

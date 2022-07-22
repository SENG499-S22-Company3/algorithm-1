package genetic

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
	"math/rand"
	"strings"

	//"fmt"
	"math"
	"strconv"

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

// Evaluate a Semester by checking for time conflicts and other properties
func (sem Semester) Evaluate() (penalty float64, err error) {
	_, _, profs, term := getInput()

	// comparing each course to each other course
	for i := range sem {
		for j := range sem {
			// ignore courses with TBD profs
			if sem[i].Prof.DisplayName == "TBD" {
				continue
			}

			// evaluate same stream time differences
			if sem[i].StreamSequence == sem[j].StreamSequence {
				if sem[i].Assignment.Monday == sem[j].Assignment.Monday {
					t1, _ := strconv.Atoi(sem[i].Assignment.BeginTime)
					t2, _ := strconv.Atoi(sem[j].Assignment.EndTime)
					diff := math.Abs(float64(t1 - t2))
					if diff > 600 {
						penalty += (diff - 600)
					}
				}
			}
		}
		// penalize low prof preference values
		for j := range sem[i].Prof.Preferences {
			if strings.Contains(sem[i].Prof.Preferences[j].CourseNum, sem[i].CourseNumber) && strings.Contains(sem[i].Prof.Preferences[j].CourseNum, sem[i].Subject) {
				penalty += 10 * (6 - float64(sem[i].Prof.Preferences[j].PreferenceNum))
			}
		}
	}

	// penalty for prof with back to back courses
	for _, p := range profs {
		var p_courses []structs.Course
		for i := range sem {
			if sem[i].Prof.DisplayName == p.DisplayName {
				p_courses = append(p_courses, sem[i])
			}
		}
		for i := range p_courses {
			t1, _ := strconv.Atoi(p_courses[i].Assignment.BeginTime)
			for j := range p_courses {
				t2, _ := strconv.Atoi(p_courses[j].Assignment.EndTime)
				if math.Abs(float64(t2-t1)) <= 10 {
					penalty += 300
				}
			}
		}
	}

	// Check if timeslots violate hard requirements
	_, fail := scheduling.BaseTimeslotMaps(sem, term)
	if fail != nil {
		penalty += 1000
	}

	// Checks if prof requirements violated
	fail = scheduling.ScheduleConstraintsCheck(term, sem, profs)
	if fail != nil {
		penalty += 1000
	}

	return
}

// Mutate a Semester by applying by permutation mutation and/or splice mutation.
func (sem Semester) Mutate(rng *rand.Rand) {
	_, _, _, term = getInput()

	if rng.Float64() < 0.45 {
		sem = scheduling.ChangeRandomCourseTime(sem, term)
		sem = scheduling.ChangeRandomCourseProf(sem)
	}
}

// Crossover a Semester with another Semester by using Partially Mixed Crossover (PMX).
func (sem Semester) Crossover(q eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNX(sem, q.(Semester), 2, rng)
}

// MakeSemester creates a random semester
func MakeSemester(rng *rand.Rand) eaopt.Genome {

	hardScheduled, coursesToSchedule, professors, term := getInput()

	testStreamtype, _ := scheduling.BaseTimeslotMaps(hardScheduled, term)
	coursesToSchedule, _, _ = scheduling.AddCoursesToStreamMaps(scheduling.Split(coursesToSchedule), testStreamtype, term)
	testScheduleCourse := scheduling.AssignCourseProf(hardScheduled, coursesToSchedule, professors, term)

	emptyMap := scheduling.CreateEmptyStreamType()
	hardCourses, _, _ := scheduling.AddCoursesToStreamMaps(hardScheduled, emptyMap, term)

	testScheduleCourse = append(testScheduleCourse, hardCourses...)

	testSem := make(Semester, len(testScheduleCourse))

	copy(testSem, testScheduleCourse)

	return testSem
}

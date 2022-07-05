package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"strconv"
)

func addSections(courses []structs.Course, course structs.Course, numSections uint, i int)([]structs.Course) {
	
	// calculate new capacity
	newCapacity := course.CourseCapacity / numSections
			
	// update original section
	courses[i].CourseCapacity = newCapacity + (course.CourseCapacity % numSections)
	courses[i].NumSections = 1

	// add new sections
	numSecCurr := uint(1)
	for numSecCurr < numSections {
		newSection := courses[i]
		newSection.CourseCapacity = newCapacity 
		newSection.SequenceNumber = newSection.SequenceNumber[:2] + strconv.Itoa(int(numSecCurr+1))
		courses = append(courses, newSection)
		numSecCurr++
	}
	return courses
}

func printResults(courses []structs.Course)(){
	for i, course := range courses {
		fmt.Println(i, course)	
	}
}

func Split(courses []structs.Course) ([]structs.Course){
    for i, course := range courses {
		numSections :=  course.NumSections
		if numSections == 1 && course.CourseCapacity > 200 {
			splitInto := uint(course.CourseCapacity / 100)
			courses = addSections(courses, course, splitInto, i)
		} else if numSections > 1 {
			courses = addSections(courses, course, numSections, i)
		}
    }
	return courses
}
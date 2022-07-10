package scheduling

import (
	"algorithm-1/structs"
	"fmt"
	"strconv"
)

func addSections(splitSchedule []structs.Course, course structs.Course, numSections uint)([]structs.Course) {
	
	// calculate new capacity
	newCapacity := course.CourseCapacity / numSections

	// update original section
	splitSchedule = append(splitSchedule, structs.Course{
		CourseNumber: course.CourseNumber,
		Subject: course.Subject,
		CourseTitle: course.CourseTitle,
		NumSections: 1,
		CourseCapacity: newCapacity + (course.CourseCapacity % numSections),
		StreamSequence: course.StreamSequence,
		SequenceNumber: "A01",
	})

	// add new sections
	numSecCurr := uint(1)
	for numSecCurr < numSections {
		splitSchedule = append(splitSchedule, structs.Course{
			CourseNumber: course.CourseNumber,
			Subject: course.Subject,
			CourseTitle: course.CourseTitle,
			NumSections: 1,
			CourseCapacity: newCapacity,
			StreamSequence: course.StreamSequence,
			SequenceNumber: "A0" + strconv.Itoa(int(numSecCurr+1)),
		})
		numSecCurr++
	}
	return splitSchedule
}

func printResults(courses []structs.Course)(){
	for i, course := range courses {
		fmt.Println(i, course)	
	}
}

func Split(courses []structs.Course) ([]structs.Course){

	splitSchedule := []structs.Course{}

    for _, course := range courses {

		numSections :=  course.NumSections
		if numSections == 0 && course.CourseCapacity > 200 {
			splitInto := uint(course.CourseCapacity / 100)
			splitSchedule = addSections(splitSchedule, course, splitInto)
		} else if numSections > 1 {
			splitSchedule = addSections(splitSchedule, course, numSections)
		} else {
			splitSchedule = append(splitSchedule, structs.Course{
				CourseNumber: course.CourseNumber,
				Subject: course.Subject,
				CourseTitle: course.CourseTitle,
				NumSections: 1,
				CourseCapacity: course.CourseCapacity,
				StreamSequence: course.StreamSequence,
				SequenceNumber: "A01",
			})
		}
    }

	return splitSchedule
}
package structs

import "time"

type Assignment struct {
	course    Course
	prof      Professor
	startDate time.Time
	endDate   time.Time
	beginTime time.Time
	endTime   time.Time
	sunday    bool
	monday    bool
	tuesday   bool
	wednesday bool
	thursday  bool
	friday    bool
	saturday  bool
}

type Course struct {
	courseNumber   uint
	subject        string
	sequenceNumber string
	courseTitle    string
}

type Professor struct {
	// Preferences in map? Maybe its own struct
	displayName string
}

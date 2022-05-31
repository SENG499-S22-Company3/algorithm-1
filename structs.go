package main

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
}

type Professor struct {
}

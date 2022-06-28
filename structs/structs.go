package structs

type Input struct {
	HardScheduled     Schedule    `json:"hardScheduled"`
	CoursesToSchedule Schedule    `json:"coursesToSchedule"`
	Professors        []Professor `json:"professors"`
}

type Schedule struct {
	FallCourses   []Course `json:"fallTermCourses"`
	SpringCourses []Course `json:"springTermCourses"`
	SummerCourses []Course `json:"summerTermCourses"`
}

type Assignment struct {
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	BeginTime string  `json:"beginTime"`
	EndTime   string  `json:"endtime"`
	HoursWeek float32 `json:"hoursWeek"`
	Sunday    bool    `json:"sunday"`
	Monday    bool    `json:"monday"`
	Tuesday   bool    `json:"tuesday"`
	Wednesday bool    `json:"wednesday"`
	Thursday  bool    `json:"thursday"`
	Friday    bool    `json:"friday"`
	Saturday  bool    `json:"saturday"`
}

type Course struct {
	CourseNumber      string     `json:"courseNumber"`
	Subject           string     `json:"subject"`
	SequenceNumber    string     `json:"sequenceNumber"`
	CourseTitle       string     `json:"courseTitle"`
	RequiredEquipment []string   `json:"requiredEquipment,omitempty"`
	StreamSequence    string     `json:"streamSequence,omitempty"`
	Assignment        Assignment `json:"meetingTime"`
	Prof              Professor  `json:"prof,omitempty"`
}

type Professor struct {
	Preferences       []Preference `json:"prefs,omitempty"`
	DisplayName       string       `json:"displayName,omitempty"`
	RequiredEquipment []string     `json:"requiredEquipment,omitempty"`
	FallTermCourses   uint         `json:"fallTermCourses,omitempty"`
	SpringTermCourses uint         `json:"springTermCourses,omitempty"`
	SummerTermCourses uint         `json:"summerTermCourses,omitempty"`
}

type Preference struct {
	CourseNum     string `json:"courseNum,omitempty"`
	PreferenceNum uint   `json:"preferenceNum,omitempty"`
	Term          string `json:"term,omitempty"`
}

// Should follow key = start time, value = list of subject and course number strings
// Ex. Calling *.Monday["1130"] returns "CSC 115"
type Timeslots struct {
	Sunday    map[string]string
	Monday    map[string]string
	Tuesday   map[string]string
	Wednesday map[string]string
	Thursday  map[string]string
	Friday    map[string]string
	Saturday  map[string]string
}

type StreamType struct {
	S1A Timeslots
	S1B Timeslots
	S2A Timeslots
	S2B Timeslots
	S3A Timeslots
	S3B Timeslots
	S4A Timeslots
	S4B Timeslots
}

// Should follow key = prof display_name, value = list of courses he's teaching
// Ex. Calling *.Monday["Zastre"] returns [ courses1, courses2 ]
type ProfAssignments struct {
	ProfAssigned map[string][]Course
}

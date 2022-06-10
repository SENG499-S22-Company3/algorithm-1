package structs

type Input struct {
	HistoricData      Schedule    `json:"historicData"`
	CoursesToSchedule Schedule    `json:"coursesToSchedule"`
	Professors        []Professor `json:"professors"`
}

type Schedule struct {
	FallCourses   []Course `json:"fallTermCourses"`
	SpringCourses []Course `json:"springTermCourses"`
	SummerCourses []Course `json:"summerTermCourses"`
}

type Assignment struct {
	StartDate string  `json:"startDate,omitempty"`
	EndDate   string  `json:"endDate,omitempty"`
	BeginTime string  `json:"beginTime,omitempty"`
	EndTime   string  `json:"endtime,omitempty"`
	HoursWeek float32 `json:"hoursWeek,omitempty"`
	Sunday    bool    `json:"sunday,omitempty"`
	Monday    bool    `json:"monday,omitempty"`
	Tuesday   bool    `json:"tuesday,omitempty"`
	Wednesday bool    `json:"wednesday,omitempty"`
	Thursday  bool    `json:"thursday,omitempty"`
	Friday    bool    `json:"friday,omitempty"`
	Saturday  bool    `json:"saturday,omitempty"`
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

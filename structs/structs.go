package structs

type Schedule struct {
	FallCourses   []Course `json:"fallTermCourses"`
	SpringCourses []Course `json:"springTermCourses"`
	SummerCourses []Course `json:"summerTermCourses"`
}

type Assignment struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	BeginTime uint   `json:"beginTime"`
	EndTime   uint   `json:"endtime"`
	HoursWeek uint   `json:"hoursWeek"`
	Sunday    bool   `json:"sunday"`
	Monday    bool   `json:"monday"`
	Tuesday   bool   `json:"tuesday"`
	Wednesday bool   `json:"wednesday"`
	Thursday  bool   `json:"thursday"`
	Friday    bool   `json:"friday"`
	Saturday  bool   `json:"saturday"`
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

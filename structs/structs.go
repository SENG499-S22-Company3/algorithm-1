package structs

type Schedule struct {
	FallCourses   []Assignment
	SpringCourses []Assignment
	SummerCourses []Assignment
}

type Assignment struct {
	Course    Course
	Prof      Professor
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	BeginTime uint   `json:"beginTime"`
	EndTime   uint   `json:"endtime"`
	Sunday    bool   `json:"sunday"`
	Monday    bool   `json:"monday"`
	Tuesday   bool   `json:"tuesday"`
	Wednesday bool   `json:"wednesday"`
	Thursday  bool   `json:"thursday"`
	Friday    bool   `json:"friday"`
	Saturday  bool   `json:"saturday"`
}

type Course struct {
	CourseNumber      uint     `json:"courseNumber"`
	Subject           string   `json:"subject"`
	SequenceNumber    string   `json:"sequenceNumber"`
	CourseTitle       string   `json:"courseTitle"`
	RequiredEquipment []string `json:"requiredEquipment,omitempty"`
	StreamSequence    string   `json:"streamSequence"`
}

type Professor struct {
	Preferences       []Preference
	DisplayName       string   `json:"displayName,omitempty"`
	RequiredEquipment []string `json:"requiredEquipment,omitempty"`
	FallTermCourses   uint     `json:"fallTermCourses,omitempty"`
	SpringTermCourses uint     `json:"springTermCourses,omitempty"`
	SummerTermCourses uint     `json:"summerTermCourses,omitempty"`
}

type Preference struct {
	CourseNum     string `json:"courseNum,omitempty"`
	PreferenceNum uint   `json:"preferenceNum,omitempty"`
	Term          string `json:"term,omitempty"`
}

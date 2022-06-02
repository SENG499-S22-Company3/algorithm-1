package structs

type Assignment struct {
	Course    Course
	Prof      Professor
	StartDate string // Follow "yyyy-mm=dd"
	EndDate   string // Follow "yyyy-mm-dd"
	BeginTime string // Use 24hr "0000" - "2359"
	EndTime   string // Use 24hr "0000" - "2359"
	Sunday    bool
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
	Saturday  bool
}

type Course struct {
	CourseNumber      uint
	Subject           string
	SequenceNumber    string
	CourseTitle       string
	RequiredEquipment []string
	RequiresPEng      bool
}

type Professor struct {
	Preferences       []Preference
	CoursesCanTeach   []Course
	DisplayName       string
	TeachingStatus    string
	RequiredEquipment []string
	HasPEng           bool
}

type Preference struct {
	Course        Course
	PreferenceNum uint
}

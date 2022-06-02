package structs

type Assignment struct {
	Course    Course
	Prof      Professor
	StartDate string
	EndDate   string
	BeginTime string
	EndTime   string
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

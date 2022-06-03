package structs

type Assignment struct {
	Course    Course
	Prof      Professor
	StartDate string // Follow "yyyy-mm=dd"
	EndDate   string // Follow "yyyy-mm-dd"
	BeginTime uint
	EndTime   uint
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
	StreamSequence    string
}

type Professor struct {
	Preferences        []Preference
	CoursesCanTeach    []Course
	DisplayName        string
	NumCoursesCanTeach uint
	RequiredEquipment  []string
	HasPEng            bool
	FallTermCourses    uint
	SpringTermCourses  uint
	SummerTermCourses  uint
}

type Preference struct {
	Course        Course
	PreferenceNum uint
	Term          string
}

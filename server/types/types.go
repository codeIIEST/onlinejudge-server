package types

// User user object
type User struct {
	Handle    string `json:"handle" form:"handle"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Rating    int32  `json:"rating" form:"rating"`
}

// Contest Table struct
type Contest struct {
	Name                string   `json:"name"`
	Type                string   `json:"type"`
	Phase               string   `json:"phase"`
	DurationSeconds     int32    `json:"durationseconds"`
	StartTimeSeconds    int32    `json:"starttimeseconds"`
	RelativeTimeSeconds int32    `json:"relativetimeseconds"`
	Author              []string `json:"author" pg:",array"`  
	Problems            []string `json:"problems" pg:",array"`
}

// Problem Table struct
type Problem struct {
	ContestID int32    `json:"contestid"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Rating    int32    `json:"rating"`
	Tags      []string `json:"tags" pg:",array"`
}

// SubmissionData Table struct
type SubmissionData struct {
	UserID          string  `json:"userid"`
	TestResultID    string  `json:"testresultid"`
	TestDataID      string  `json:"testdataid"`
	ContestID       string  `json:"contestid"`
	ProblemID       string  `json:"problemid"`
	CreationTime    string  `json:"creationtime"`
	RelativeTime    string  `json:"relativetime"`
	Verdict         string  `json:"verdict"`
	Language        string  `json:"language"`
	PassedTestCount int32   `json:"passedtestcount"`
	TimeConsumed    float64 `json:"timeconsumed"`
	MemoryConsumed  int64   `json:"memoryconsumed"`
	Point           float64 `json:"point"`
}

// TestResult Table struct
type TestResult struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Time    []float64 `json:"time"`
	Memory  []float64 `json:"memory" pg:",array"`
	Result  []string  `json:"result" pg:",array"`
	Error   []string  `json:"error" pg:",array"`
}

// TestData Table struct
type TestData struct {
	Lang       string   `json:"lang"`
	Filename   string   `json:"filename"`
	Code       string   `json:"code"`
	Path       string   `json:"path"`
	Image      string   `json:"image"`
	TestCount  int      `json:"testcount"`
	InputData  []string `json:"inputdata" pg:",array"`
	OutputData []string `json:"outputdata" pg:",array"`
	TimeLimit  int      `json:"timelimit"`
	MemLimit   int64    `json:"memlimit"`
}

package handler

type Database struct {
	programs map[string]Program
}

type Program struct {
	programName        string
	programDescription string
	courses            map[string]Course
	plos               map[string]PLO
}

type PLO struct {
	ploName        string
	ploDescription string
}

type Course struct {
	courseName        string
	courseDescription string
	semester          int
	year              int
	students          map[string]Student
	los               map[string]LO
	quizzes           map[string]Quiz
}

type Student struct {
	studentEmail   string
	studentName    string
	studentSurname string
}

type LO struct {
	loTitle      string
	levels       []LOLevel
	linkedploIDs map[string]bool
}

type LOLevel struct {
	level            int
	levelDescription string
}

type Quiz struct {
	quizName  string
	questions map[string]Question
}

type Question struct {
	questionTitle string
	maxScore      int
	results       []QuestionResult
	linkedloIDs   map[string]int
}

type QuestionResult struct {
	studentID    string
	studentScore int
}

type QuestionExcel struct {
	QuestionTitle string `json:"questionTitle"`
	Maxscore      int    `json:"maxScore"`
	StudentID     int    `json:"studentID"`
	StudentScore  int    `json:"studentScore"`
}

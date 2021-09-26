package course

import (
	"sync"

	"github.com/google/uuid"
)

type (
	Course struct {
		CourseID   string `json:"courseID"`
		ProgramID  string `json:"programID"`
		CourseName string `json:"courseName"`
		Semester   int    `json:"semester"`
		Year       int    `json:"year"`
	}
	Courses []Course
)

var (
	once    sync.Once
	courses *Courses
)

func Connect() {
	once.Do(func() {
		courses = new(Courses)
	})
}

func AddRecord(programID, name string, semester, year int) {
	*courses = append(*courses, Course{uuid.New().String(), programID, name, semester, year})
}

func GetRecords(programID string) Courses {
	response := make(Courses, 0)
	for _, course := range *courses {
		if course.ProgramID == programID {
			response = append(response, course)
		}
	}
	return response
}

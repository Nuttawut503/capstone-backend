package student

import "sync"

type (
	Student struct {
		CourseID    string `json:"courseID"`
		StudentID   int    `json:"studentID"`
		StudentName string `json:"studentName"`
	}
	Students []Student
)

var (
	once     sync.Once
	students *Students
)

func Connect() {
	once.Do(func() {
		students = new(Students)
	})
}

func GetRecords(courseID string) Students {
	response := make(Students, 0)
	for _, student := range *students {
		if student.CourseID == courseID {
			response = append(response, student)
		}
	}
	return response
}

func AddRecords(newStudents Students) {
	*students = append(*students, newStudents...)
}

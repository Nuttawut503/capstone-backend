package quiz

import (
	"sync"

	"github.com/google/uuid"
)

type (
	Quiz struct {
		QuizID   string `json:"quizID"`
		courseID string
		Title    string `json:"title"`
	}
	Quizzes []Quiz
)

var (
	once    sync.Once
	quizzes *Quizzes
)

func Connect() {
	once.Do(func() {
		quizzes = new(Quizzes)
	})
}

func AddRecord(courseID, title string) {
	*quizzes = append(*quizzes, Quiz{uuid.New().String(), courseID, title})
}

func GetQuizzesByCourseID(courseID string) Quizzes {
	var response Quizzes
	for _, quiz := range *quizzes {
		if quiz.courseID == courseID {
			response = append(response, quiz)
		}
	}
	return response
}

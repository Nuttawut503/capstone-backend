package question

import (
	"sync"

	"github.com/google/uuid"
)

type (
	Question struct {
		questionID, quizID, title string
		maxScore                  int
	}
	Questions []Question
)

var (
	once      sync.Once
	questions *Questions
)

func Connect() {
	once.Do(func() {
		questions = new(Questions)
	})
}

func AddRecord(quizID, title string, maxScore int) string {
	questionID := uuid.New().String()
	*questions = append(*questions, Question{questionID, quizID, title, maxScore})
	return questionID
}

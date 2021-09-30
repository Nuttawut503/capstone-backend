package question

import (
	"sync"

	"github.com/google/uuid"
)

type (
	Question struct {
		quizID     string
		QuestionID string `json:"questionID"`
		Title      string `json:"title"`
		MaxScore   int    `json:"maxScore"`
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
	*questions = append(*questions, Question{quizID, questionID, title, maxScore})
	return questionID
}

func GetQuestionsByQuizID(quizID string) Questions {
	var response Questions
	for _, question := range *questions {
		if question.quizID == quizID {
			response = append(response, question)
		}
	}
	return response
}

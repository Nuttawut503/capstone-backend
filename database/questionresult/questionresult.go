package questionresult

import "sync"

type (
	QuestionResult struct {
		questionID string
		studentID  int
		score      int
	}
	QuestionResults []QuestionResult
)

var (
	once            sync.Once
	questionresults *QuestionResults
)

func Connect() {
	once.Do(func() {
		questionresults = new(QuestionResults)
	})
}

func AddRecord(questionID string, studentID int, score int) {
	*questionresults = append(*questionresults, QuestionResult{questionID, studentID, score})
}

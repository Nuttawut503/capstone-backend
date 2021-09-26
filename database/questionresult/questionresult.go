package questionresult

import "sync"

type (
	QuestionResult struct {
		questionID, studentID string
		score                 int
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

func AddRecord(questionID, studentID string, score int) {
	*questionresults = append(*questionresults, QuestionResult{questionID, studentID, score})
}

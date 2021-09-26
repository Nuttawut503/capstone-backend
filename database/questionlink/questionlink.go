package questionlink

import "sync"

type (
	QuestionLink struct {
		questionID, loID string
		loLevel          int
	}
	QuestionLinks []QuestionLink
)

var (
	once          sync.Once
	questionlinks *QuestionLinks
)

func Connect() {
	once.Do(func() {
		questionlinks = new(QuestionLinks)
	})
}

func AddRecord(questionID, loID string, loLevel int) {
	*questionlinks = append(*questionlinks, QuestionLink{questionID, loID, loLevel})
}

func GetAllLOLinksByQuestionID(questionID string) {

}

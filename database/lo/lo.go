package lo

import (
	"sync"

	"github.com/google/uuid"
)

type (
	LO struct {
		loID, courseID, info string
	}
	LOs []LO
)

var (
	once sync.Once
	los  *LOs
)

func Connect() {
	once.Do(func() {
		los = new(LOs)
	})
}

func AddRecord(courseID, info string) {
	*los = append(*los, LO{uuid.New().String(), courseID, info})
}

func GetLOsByCourseID(courseID string) [][2]string {
	response := make([][2]string, 0)
	for _, lo := range *los {
		if lo.courseID == courseID {
			response = append(response, [2]string{lo.loID, lo.info})
		}
	}
	return response
}

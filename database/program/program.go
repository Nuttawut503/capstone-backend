package program

import (
	"sync"

	"github.com/google/uuid"
)

type (
	Program struct {
		ProgramID   string `json:"programID"`
		UserID      string `json:"userID"`
		ProgramName string `json:"programName"`
		Description string `json:"description"`
	}
	Programs []Program
)

var (
	once     sync.Once
	programs *Programs
)

func Connect() {
	once.Do(func() {
		programs = new(Programs)
	})
}

func AddRecord(programName string) {
	*programs = append(*programs, Program{uuid.New().String(), "1234", programName, ""})
}

func GetRecords() Programs {
	return *programs
}

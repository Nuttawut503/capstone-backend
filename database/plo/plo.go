package plo

import (
	"sync"

	"github.com/google/uuid"
)

type (
	PLO struct {
		ploID, programID string
	}
	PLOs []PLO
)

var (
	once sync.Once
	plos *PLOs
)

func Connect() {
	once.Do(func() {
		plos = new(PLOs)
	})
}
func AddRecord(programID string) string {
	ploID := uuid.New().String()
	*plos = append(*plos, PLO{ploID, programID})
	return ploID
}

func GetPLOIDsByProgranID(programID string) []string {
	ploIDs := make([]string, 0)
	for _, plo := range *plos {
		if plo.programID == programID {
			ploIDs = append(ploIDs, plo.ploID)
		}
	}
	return ploIDs
}

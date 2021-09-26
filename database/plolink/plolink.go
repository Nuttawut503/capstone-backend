package plolink

import (
	"sync"
)

type (
	PLOLink struct {
		ploID, loID string
	}
	PLOLinks []PLOLink
)

var (
	once     sync.Once
	plolinks *PLOLinks
)

func Connect() {
	once.Do(func() {
		plolinks = new(PLOLinks)
	})
}
func AddRecord(ploID, loID string) {
	*plolinks = append(*plolinks, PLOLink{ploID, loID})
}

package plolink

import (
	"sync"
)

type (
	PLOLink struct {
		PLOID, LOID string
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

func GetPLOLinksByLOID(loID string) PLOLinks {
	response := make(PLOLinks, 0)
	for _, plolink := range *plolinks {
		if plolink.LOID == loID {
			response = append(response, plolink)
		}
	}
	return response
}

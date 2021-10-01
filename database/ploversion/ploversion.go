package ploversion

import (
	"backend/database/plolink"
	"sync"
)

type (
	PLOVersion struct {
		PloID   string `json:"ploID"`
		Info    string `json:"info"`
		version int
	}
	PLOVersions []PLOVersion
)

var (
	once        sync.Once
	ploversions *PLOVersions
)

func Connect() {
	once.Do(func() {
		ploversions = new(PLOVersions)
	})
}

func AddRecord(ploID, info string, version int) {
	*ploversions = append(*ploversions, PLOVersion{ploID, info, version})
}

func GetRecords(ploIDs []string) PLOVersions {
	response := make(PLOVersions, 0)
	for _, plo := range *ploversions {
		// if plo.PloID in ploIDs; append
		for _, ploID := range ploIDs {
			if plo.PloID == ploID {
				response = append(response, plo)
			}
		}
	}
	return response
}

func GetPLODetailByPLOIDs(pls plolink.PLOLinks) PLOVersions {
	response := make(PLOVersions, 0)
	for _, plo := range *ploversions {
		for _, pl := range pls {
			if plo.PloID == pl.PLOID {
				response = append(response, plo)
			}
		}
	}
	return response
}

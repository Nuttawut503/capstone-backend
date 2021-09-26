package lolevel

import "sync"

type (
	LOLevel struct {
		loID, info string
		level      int
	}
	LOLevels []LOLevel
)

var (
	once     sync.Once
	lolevels *LOLevels
)

func Connect() {
	once.Do(func() {
		lolevels = new(LOLevels)
	})
}

func AddRecord(loID, info string, level int) {
	*lolevels = append(*lolevels, LOLevel{loID, info, level})
}

func GetLevelInfoByLOID(loID string) (response [][2]interface{}) {
	for _, lolevel := range *lolevels {
		if lolevel.loID == loID {
			response = append(response, [2]interface{}{lolevel.level, lolevel.info})
		}
	}
	return response
}

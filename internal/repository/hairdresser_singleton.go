package repository

import "sync"

var hairdresserLock = &sync.Mutex{}
var hairdresserRepo *Hairdresser

func GetHairdresserRepository() *Hairdresser {
	if hairdresserRepo == nil {
		hairdresserLock.Lock()
		defer hairdresserLock.Unlock()

		if hairdresserRepo == nil {
			hairdresserRepo = newHairdresser(GetDB())
			return hairdresserRepo
		} else {
			return hairdresserRepo
		}
	} else {
		return hairdresserRepo
	}
}

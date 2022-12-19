package repository

import "sync"

var serviceInfoLock = &sync.Mutex{}
var serviceInfoRepo *ServiceInfo

func GetServiceInfoRepository() *ServiceInfo {
	if serviceInfoRepo == nil {
		serviceInfoLock.Lock()
		defer serviceInfoLock.Unlock()

		if serviceInfoRepo == nil {
			serviceInfoRepo = newServiceInfo(GetDB())
			return serviceInfoRepo
		} else {
			return serviceInfoRepo
		}
	} else {
		return serviceInfoRepo
	}
}

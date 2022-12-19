package service

import (
	"net/http"
	"sync"

	"github.com/ew0s/kontora-autotests/swagger"
)

var lock = &sync.Mutex{}

var singleKontoraClient *swagger.APIClient

func GetKontoraClient() *swagger.APIClient {
	if singleKontoraClient == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleKontoraClient == nil {
			cfg := &swagger.Configuration{
				BasePath:   "http://localhost:8082",
				HTTPClient: http.DefaultClient,
			}

			singleKontoraClient = swagger.NewAPIClient(cfg)

			return singleKontoraClient
		} else {
			return singleKontoraClient
		}
	} else {
		return singleKontoraClient
	}
}

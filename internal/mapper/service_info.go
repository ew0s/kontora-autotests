package mapper

import (
	"github.com/ew0s/kontora-autotests/internal/domain/entities"
	"github.com/ew0s/kontora-autotests/swagger"
)

func MakeServiceInfoEntity(swaggerDTO swagger.ServiceInfoDto) entities.ServiceInfo {
	var hairdresserID int64
	if swaggerDTO.HairdresserDto != nil {
		hairdresserID = swaggerDTO.HairdresserDto.Id
	}

	var serviceID int64
	if swaggerDTO.ServiceDto != nil {
		serviceID = swaggerDTO.ServiceDto.Id
	}

	return entities.ServiceInfo{
		HairdresserID: hairdresserID,
		Id:            swaggerDTO.Id,
		Price:         swaggerDTO.Price,
		ServiceID:     serviceID,
	}
}

func MakeServiceInfoEntities(swaggerDTOs []swagger.ServiceInfoDto) []entities.ServiceInfo {
	res := make([]entities.ServiceInfo, 0, len(swaggerDTOs))

	for _, dto := range swaggerDTOs {
		res = append(res, MakeServiceInfoEntity(dto))
	}

	return res
}

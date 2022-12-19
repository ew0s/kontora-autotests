package mapper

import (
	"github.com/ew0s/kontora-autotests/internal/domain/entities"
	"github.com/ew0s/kontora-autotests/swagger"
)

func MakeHairdresserEntity(swaggerDTO swagger.HairdresserDto) entities.Hairdresser {
	return entities.Hairdresser{
		Description: swaggerDTO.Description,
		FullName:    swaggerDTO.FullName,
		Id:          swaggerDTO.Id,
		Phone:       swaggerDTO.Phone,
	}
}

func MakeHairdresserEntities(swaggerDTOs []swagger.HairdresserDto) []entities.Hairdresser {
	res := make([]entities.Hairdresser, 0, len(swaggerDTOs))

	for _, dto := range swaggerDTOs {
		res = append(res, MakeHairdresserEntity(dto))
	}

	return res
}

package mapper

import (
	"github.com/ew0s/kontora-autotests/internal/domain/entities"
	"github.com/ew0s/kontora-autotests/swagger"
)

func MakeServiceEntity(swaggerDTO swagger.ServiceDto) entities.Service {
	return entities.Service{
		Id:   swaggerDTO.Id,
		Name: swaggerDTO.Name,
	}
}

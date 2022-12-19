package repository

import (
	"context"

	"github.com/ew0s/kontora-autotests/internal/domain/entities"
)

type ServiceInfo interface {
	Truncate(ctx context.Context) error
	InsertServiceInfos(ctx context.Context, serviceInfos []entities.ServiceInfo) error
}

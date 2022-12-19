package repository

import (
	"context"

	"github.com/ew0s/kontora-autotests/internal/domain/entities"
)

type Hairdresser interface {
	Truncate(ctx context.Context) error
	InsertHairdressers(ctx context.Context, hairdressers []entities.Hairdresser) error
}

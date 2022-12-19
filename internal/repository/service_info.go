package repository

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/ew0s/kontora-autotests/internal/domain/entities"
	"github.com/ew0s/kontora-autotests/internal/repository/schema"
	"github.com/jmoiron/sqlx"
)

type ServiceInfo struct {
	db *sqlx.DB
}

func newServiceInfo(db *sqlx.DB) *ServiceInfo {
	return &ServiceInfo{db: db}
}

func (r *ServiceInfo) Truncate(ctx context.Context) error {
	query := goqu.New("postgres", r.db).
		Truncate(goqu.T(string(schema.ServiceInfo))).Identity("RESTART").Cascade().
		Prepared(true).
		Executor()

	if _, err := query.ExecContext(ctx); err != nil {
		return fmt.Errorf("executing query with context: %w", err)
	}

	return nil
}

func (r *ServiceInfo) InsertServiceInfos(ctx context.Context, serviceInfos []entities.ServiceInfo) error {
	query := goqu.New("postgres", r.db).
		Insert(goqu.T(string(schema.ServiceInfo))).
		Rows(serviceInfos).
		Prepared(true).
		Executor()

	if _, err := query.ExecContext(ctx); err != nil {
		return fmt.Errorf("executing query with context: %w", err)
	}

	return nil
}

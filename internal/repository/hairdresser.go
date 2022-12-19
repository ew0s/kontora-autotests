package repository

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/ew0s/kontora-autotests/internal/domain/entities"
	"github.com/ew0s/kontora-autotests/internal/repository/schema"
	"github.com/jmoiron/sqlx"
)

type Hairdresser struct {
	db *sqlx.DB
}

func newHairdresser(db *sqlx.DB) *Hairdresser {
	return &Hairdresser{db: db}
}

func (r *Hairdresser) Truncate(ctx context.Context) error {
	query := goqu.New("postgres", r.db).
		Truncate(goqu.T(string(schema.Hairdresser))).Identity("RESTART").Cascade().
		Prepared(true).
		Executor()

	if _, err := query.ExecContext(ctx); err != nil {
		return fmt.Errorf("executing query with context: %w", err)
	}

	return nil
}

func (r *Hairdresser) InsertHairdressers(ctx context.Context, hairdressers []entities.Hairdresser) error {
	query := goqu.New("postgres", r.db).
		Insert(goqu.T(string(schema.Hairdresser))).
		Rows(hairdressers).
		Prepared(true).
		Executor()

	if _, err := query.ExecContext(ctx); err != nil {
		return fmt.Errorf("executing query with context: %w", err)
	}

	return nil
}

package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loanem-backend/participant-service/infra/database/sqlc"
	"github.com/loanem-backend/participant-service/internal/entity"
)

type TeamRepository interface {
	WithTX(tx pgx.Tx) TeamRepository

	Insert(ctx context.Context, t *entity.Team) error
}

type teamRepository struct {
	db *sqlc.Queries
}

func NewTeamRepository(q *sqlc.Queries) TeamRepository {
	return &teamRepository{
		db: q,
	}
}

func (r *teamRepository) WithTX(tx pgx.Tx) TeamRepository {
	return &teamRepository{
		db: r.db.WithTx(tx),
	}
}

func (r *teamRepository) Insert(ctx context.Context, t *entity.Team) error {
	if err := r.db.InsertTeam(ctx, sqlc.InsertTeamParams{
		ID:      t.ID.String(),
		Number:  int16(t.Number),
		ClassID: pgtype.Int4{Int32: int32(t.Class.ID), Valid: true},
	}); err != nil {
		return err
	}

	return nil
}

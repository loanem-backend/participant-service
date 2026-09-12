package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loanem-backend/participant-service/infra/database/sqlc"
	"github.com/loanem-backend/participant-service/internal/entity"
)

type ParticipantRepository interface {
	Insert(ctx context.Context, p *entity.Participant) error
}

type participantRepository struct {
	db *sqlc.Queries
}

func NewParticipantRepository(q *sqlc.Queries) ParticipantRepository {
	return &participantRepository{
		db: q,
	}
}

func (r *participantRepository) Insert(ctx context.Context, p *entity.Participant) error {
	if err := r.db.InsertParticipant(ctx, sqlc.InsertParticipantParams{
		ID:      p.ID.String(),
		Nim:     p.Nim,
		Name:    p.Name,
		ClassID: pgtype.Int4{Int32: int32(p.Class.ID), Valid: true},
	}); err != nil {
		return err
	}

	return nil
}

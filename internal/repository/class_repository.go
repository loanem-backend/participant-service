package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loanem-backend/participant-service/infra/database/sqlc"
	"github.com/loanem-backend/participant-service/internal/entity"
)

type ClassRepository interface {
	Insert(ctx context.Context, c *entity.Class) (int32, error)
	Delete(ctx context.Context, cID int32) error
	FindByCourse(ctx context.Context, crsID int32) ([]*entity.Class, error)
}

type classRepository struct {
	db *sqlc.Queries
}

func NewClassRepository(q *sqlc.Queries) ClassRepository {
	return &classRepository{
		db: q,
	}
}

func (r *classRepository) Insert(ctx context.Context, c *entity.Class) (int32, error) {
	result, err := r.db.InsertClass(ctx, sqlc.InsertClassParams{
		Name:     c.Name,
		CourseID: pgtype.Int4{Int32: int32(c.Course.ID), Valid: true},
	})
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *classRepository) Delete(ctx context.Context, cID int32) error {
	if err := r.db.DeleteClass(ctx, cID); err != nil {
		return err
	}

	return nil
}

func (r *classRepository) FindByCourse(ctx context.Context, crsID int32) ([]*entity.Class, error) {
	rows, err := r.db.FindClassesByCourseID(ctx, pgtype.Int4{
		Int32: crsID, Valid: true,
	})
	if err != nil {
		return nil, err
	}

	classes := make([]*entity.Class, len(rows))
	for i, row := range rows {
		classes[i] = toClass(row)
	}

	return classes, nil
}

func toClass(row sqlc.FindClassesByCourseIDRow) *entity.Class {
	return &entity.Class{
		ID:   int(row.ID),
		Name: row.Name,
		Course: entity.Course{
			ID:   int(row.CourseID.Int32),
			Name: row.CourseName,
		},
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}

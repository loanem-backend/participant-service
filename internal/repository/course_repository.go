package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loanem-backend/participant-service/infra/database/sqlc"
	"github.com/loanem-backend/participant-service/internal/entity"
)

type CourseRepository interface {
	Insert(ctx context.Context, c *entity.Course) error
	Delete(ctx context.Context, cID int32) error
	UpdateName(ctx context.Context, c *entity.Course) error
}

type courseRepository struct {
	db *sqlc.Queries
}

func NewCourseRepository(q *sqlc.Queries) CourseRepository {
	return &courseRepository{
		db: q,
	}
}

func (r *courseRepository) Insert(ctx context.Context, c *entity.Course) error {
	if err := r.db.InsertCourse(ctx, sqlc.InsertCourseParams{
		ID:   int32(c.ID),
		Name: c.Name,
		Year: int32(c.Year),
	}); err != nil {
		return err
	}

	return nil
}

func (r *courseRepository) Delete(ctx context.Context, cID int32) error {
	if err := r.db.DeleteCourseByID(ctx, cID); err != nil {
		return err
	}

	return nil
}

func (r *courseRepository) UpdateName(ctx context.Context, c *entity.Course) error {
	if err := r.db.UpdateCourseName(ctx, sqlc.UpdateCourseNameParams{
		ID:        int32(c.ID),
		Name:      c.Name,
		UpdatedAt: pgtype.Timestamp{Time: c.UpdatedAt, Valid: true},
	}); err != nil {
		return err
	}

	return nil
}

package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/loanem-backend/course-service/pkg/dbtx"
	"github.com/loanem-backend/participant-service/internal/entity"
	"github.com/loanem-backend/participant-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClassService interface {
	AddClasses(ctx context.Context, courseID int32, names []string) error
}

type classService struct {
	db        *pgxpool.Pool
	classRepo repository.ClassRepository
}

func NewClassService(p *pgxpool.Pool, cr repository.ClassRepository) ClassService {
	return &classService{
		db:        p,
		classRepo: cr,
	}
}

func (s *classService) AddClasses(ctx context.Context, courseID int32, names []string) error {
	tx, err := dbtx.BeginTransaction(ctx, s.db)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	defer func() {
		err = dbtx.CommitOrRollbackTransaction(ctx, tx, err)
	}()

	for _, name := range names {
		if _, err := s.classRepo.WithTX(tx).Insert(ctx, &entity.Class{
			Name: name,
			Course: entity.Course{
				ID: int(courseID),
			},
		}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/loanem-backend/course-service/pkg/dbtx"
	"github.com/loanem-backend/participant-service/internal/entity"
	"github.com/loanem-backend/participant-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TeamService interface {
	Add(ctx context.Context, classID, numberOfTeams int32) error
}

type teamService struct {
	db       *pgxpool.Pool
	teamRepo repository.TeamRepository
}

func NewTeamService(p *pgxpool.Pool, tr repository.TeamRepository) TeamService {
	return &teamService{
		db:       p,
		teamRepo: tr,
	}
}

func (s *teamService) Add(ctx context.Context, classID, numberOfTeams int32) error {
	tx, err := dbtx.BeginTransaction(ctx, s.db)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	defer func() {
		err = dbtx.CommitOrRollbackTransaction(ctx, tx, err)
	}()

	for i := 1; i <= int(numberOfTeams); i++ {
		if err := s.teamRepo.WithTX(tx).Insert(ctx, &entity.Team{
			ID:     uuid.New(),
			Number: i,
			Class: entity.Class{
				ID: int(classID),
			},
		}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

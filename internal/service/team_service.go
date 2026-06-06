package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/loanem-backend/participant-service/internal/entity"
	"github.com/loanem-backend/participant-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TeamService interface {
	Add(ctx context.Context, t *entity.Team) (string, error)
}

type teamService struct {
	teamRepo repository.TeamRepository
}

func NewTeamService(tr repository.TeamRepository) TeamService {
	return &teamService{
		teamRepo: tr,
	}
}

func (s *teamService) Add(ctx context.Context, t *entity.Team) (string, error) {
	t.ID = uuid.New()

	if err := s.teamRepo.Insert(ctx, t); err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}

	return t.ID.String(), nil
}

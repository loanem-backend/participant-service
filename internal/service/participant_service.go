package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/loanem-backend/participant-service/internal/entity"
	"github.com/loanem-backend/participant-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ParticipantService interface {
	Create(ctx context.Context, p *entity.Participant) (string, error)
}

type participantService struct {
	participantRepo repository.ParticipantRepository
}

func NewParticipantService(pr repository.ParticipantRepository) ParticipantService {
	return &participantService{
		participantRepo: pr,
	}
}

func (s *participantService) Create(ctx context.Context, p *entity.Participant) (string, error) {
	p.ID = uuid.New()

	if err := s.participantRepo.Insert(ctx, p); err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}

	return p.ID.String(), nil
}

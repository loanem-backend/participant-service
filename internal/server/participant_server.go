package server

import (
	"context"

	"github.com/loanem-backend/participant-service/internal/mapper"
	"github.com/loanem-backend/participant-service/internal/service"
	pbparticipant "github.com/loanem-backend/protos/pb/proto/services/participant/v1"
)

type ParticipantServer struct {
	pbparticipant.UnimplementedParticipantServiceServer
	participantServ service.ParticipantService
}

func NewParticipantServer(ps service.ParticipantService) *ParticipantServer {
	return &ParticipantServer{
		participantServ: ps,
	}
}

func (s *ParticipantServer) AddParticipant(ctx context.Context, req *pbparticipant.AddParticipantRequest) (*pbparticipant.AddParticipantResponse, error) {
	idData, err := s.participantServ.Create(ctx, mapper.AddParticipantRequestToParticipant(req))
	if err != nil {
		return nil, err
	}

	return mapper.StringToAddParticipantResponse(idData), nil
}

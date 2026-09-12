package server

import (
	"github.com/loanem-backend/participant-service/internal/service"
	pbparticipant "github.com/loanem-backend/protos/pb/proto/services/participant/v1"
	"google.golang.org/grpc"
)

func Start(s *grpc.Server, cs service.ClassService, ts service.TeamService, ps service.ParticipantService) {
	pbparticipant.RegisterTeamServiceServer(s, NewTeamServer(cs, ts))
	pbparticipant.RegisterParticipantServiceServer(s, NewParticipantServer(ps))
}

package server

import (
	"context"

	"github.com/loanem-backend/participant-service/internal/service"
	pbparticipant "github.com/loanem-backend/protos/pb/proto/services/participant/v1"
)

type TeamServer struct {
	pbparticipant.UnimplementedTeamServiceServer
	classServ service.ClassService
	teamServ  service.TeamService
}

func NewTeamServer(cs service.ClassService, ts service.TeamService) *TeamServer {
	return &TeamServer{
		classServ: cs,
		teamServ:  ts,
	}
}

func (s *TeamServer) AddClass(ctx context.Context, req *pbparticipant.AddClassRequest) (*pbparticipant.AddClassResponse, error) {
	return nil, nil
}

func (s *TeamServer) AddTeam(ctx context.Context, req *pbparticipant.AddTeamRequest) (*pbparticipant.AddTeamResponse, error) {
	return nil, nil
}

func (s *TeamServer) GetClassesByCourseID(ctx context.Context, req *pbparticipant.GetClassesByCourseIDRequest) (*pbparticipant.GetClassesByCourseIDResponse, error) {
	return nil, nil
}

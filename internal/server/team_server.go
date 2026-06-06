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

func (s *TeamServer) AddClasses(ctx context.Context, req *pbparticipant.AddClassesRequest) (*pbparticipant.AddClassesResponse, error) {
	if err := s.classServ.AddClasses(ctx, req.GetCourseId(), req.GetNames()); err != nil {
		return nil, err
	}

	return &pbparticipant.AddClassesResponse{}, nil
}

func (s *TeamServer) AddTeam(ctx context.Context, req *pbparticipant.AddTeamsRequest) (*pbparticipant.AddTeamsResponse, error) {
	if err := s.teamServ.Add(ctx, req.GetClassId(), req.GetNumberOfTeams()); err != nil {
		return nil, err
	}

	return &pbparticipant.AddTeamsResponse{}, nil
}

func (s *TeamServer) GetClassesByCourseID(ctx context.Context, req *pbparticipant.GetClassesByCourseIDRequest) (*pbparticipant.GetClassesByCourseIDResponse, error) {
	return nil, nil
}

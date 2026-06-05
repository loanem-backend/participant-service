package server

import "github.com/loanem-backend/participant-service/internal/service"

type TeamServer struct {
	classServ service.ClassService
	teamServ  service.TeamService
}

func NewTeamServer(cs service.ClassService, ts service.TeamService) *TeamServer {
	return &TeamServer{
		classServ: cs,
		teamServ:  ts,
	}
}

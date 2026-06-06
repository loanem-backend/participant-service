package mapper

import (
	"github.com/loanem-backend/participant-service/internal/entity"
	pbparticipant "github.com/loanem-backend/protos/pb/proto/services/participant/v1"
)

func AddTeamRequestToTeam(req *pbparticipant.AddTeamsRequest) *entity.Team {
	return &entity.Team{
		Number: int(req.GetNumber()[0]),
		Class: entity.Class{
			ID: int(req.GetClassId()),
		},
	}
}

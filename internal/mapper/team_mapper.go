package mapper

import (
	"github.com/loanem-backend/participant-service/internal/entity"
	pbparticipant "github.com/loanem-backend/protos/pb/proto/services/participant/v1"
)

func AddTeamRequestToTeam(req *pbparticipant.AddTeamRequest) *entity.Team {
	return &entity.Team{
		Number: int(req.GetNumber()),
		Class: entity.Class{
			ID: int(req.GetClassId()),
		},
	}
}

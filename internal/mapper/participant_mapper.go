package mapper

import (
	"github.com/loanem-backend/participant-service/internal/entity"
	pbparticipant "github.com/loanem-backend/protos/pb/proto/services/participant/v1"
)

func AddParticipantRequestToParticipant(req *pbparticipant.AddParticipantRequest) *entity.Participant {
	return &entity.Participant{
		Nim:  req.GetNim(),
		Name: req.GetName(),
		Class: entity.Class{
			ID: int(req.GetClassId()),
		},
	}
}

func StringToAddParticipantResponse(s string) *pbparticipant.AddParticipantResponse {
	return &pbparticipant.AddParticipantResponse{
		Id: s,
	}
}

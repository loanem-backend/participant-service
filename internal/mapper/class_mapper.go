package mapper

import (
	"github.com/loanem-backend/participant-service/internal/entity"
	pbcourse "github.com/loanem-backend/protos/pb/proto/services/course/v1"
	pbparticipant "github.com/loanem-backend/protos/pb/proto/services/participant/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ClassToPBClass(c *entity.Class) *pbparticipant.Class {
	return &pbparticipant.Class{
		Id:   int32(c.ID),
		Name: c.Name,
		Course: &pbcourse.Course{
			Id:   int32(c.Course.ID),
			Name: c.Course.Name,
			Year: int32(c.Course.Year),
		},
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(c.UpdatedAt),
	}
}

func ClassesToGetClassesByCourseIDResponse(classes []*entity.Class) *pbparticipant.GetClassesByCourseIDResponse {
	pbClasses := make([]*pbparticipant.Class, len(classes))

	for i, c := range classes {
		pbClasses[i] = ClassToPBClass(c)
	}

	return &pbparticipant.GetClassesByCourseIDResponse{
		Classes: pbClasses,
	}
}

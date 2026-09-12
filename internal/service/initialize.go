package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/loanem-backend/participant-service/infra/database/sqlc"
	"github.com/loanem-backend/participant-service/internal/repository"
)

func Initialize(p *pgxpool.Pool) (CourseService, ClassService, TeamService, ParticipantService) {
	queries := sqlc.New(p)

	var (
		courseRepo      = repository.NewCourseRepository(queries)
		classRepo       = repository.NewClassRepository(queries)
		teamRepo        = repository.NewTeamRepository(queries)
		participantRepo = repository.NewParticipantRepository(queries)
	)

	var (
		courseServ      = NewCourseService(courseRepo)
		classServ       = NewClassService(p, classRepo)
		teamServ        = NewTeamService(p, teamRepo)
		participantServ = NewParticipantService(participantRepo)
	)

	return courseServ, classServ, teamServ, participantServ
}

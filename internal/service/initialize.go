package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/loanem-backend/participant-service/infra/database/sqlc"
	"github.com/loanem-backend/participant-service/internal/repository"
)

func Initialize(p *pgxpool.Pool) ClassService {
	queries := sqlc.New(p)

	var (
		courseRepo = repository.NewCourseRepository(queries)
		classRepo  = repository.NewClassRepository(queries)
	)

	var (
		classServ = NewClassService()
	)

	return classServ
}

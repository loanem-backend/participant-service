package service

type TeamService interface {
}

type teamService struct {
}

func NewTeamService() TeamService {
	return &teamService{}
}

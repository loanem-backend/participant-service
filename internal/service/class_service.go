package service

type ClassService interface {
}

type classService struct {
}

func NewClassService() ClassService {
	return &classService{}
}

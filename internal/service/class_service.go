package service

import "github.com/loanem-backend/participant-service/internal/repository"

type ClassService interface {
}

type classService struct {
	classRepo repository.ClassRepository
}

func NewClassService(cr repository.ClassRepository) ClassService {
	return &classService{
		classRepo: cr,
	}
}

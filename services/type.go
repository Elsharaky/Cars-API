package services

import (
	"github.com/Elsharaky/Cars-API.git/models"
	"github.com/Elsharaky/Cars-API.git/repositories"
)

type TypeService interface {
	CreateType(t *models.Type) error
	GetTypes() ([]*models.Type, error)
}

type typeService struct {
	repo repositories.TypeRepository
}

func NewTypeService() *typeService {
	return &typeService{
		repo: repositories.NewTypeRepository(),
	}
}

func (ts *typeService) CreateType(t *models.Type) error {
	return ts.repo.CreateType(t)
}

func (ts *typeService) GetTypes() ([]*models.Type, error) {
	return ts.repo.GetTypes()
}

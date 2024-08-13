package services

import (
	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/domain/models"
	"ngodeyuk-core/internal/domain/repositories"
)

type UnitService interface {
	Create(dto *dtos.UnitDTO) error
	GetAll() ([]models.Unit, error)
	GetByID(unitId uint) (*models.Unit, error)
}

type unitService struct {
	repository repositories.UnitRepository
}

func NewUnitService(repository repositories.UnitRepository) UnitService {
	return &unitService{repository}
}

func (service *unitService) Create(dto *dtos.UnitDTO) error {
	unit := &models.Unit{
		Title:       dto.Title,
		Description: dto.Description,
		Sequence:    dto.Sequence,
	}
	return service.repository.Create(unit)
}

func (service *unitService) GetAll() ([]models.Unit, error) {
	return service.repository.FindAll()
}

func (service *unitService) GetByID(unitId uint) (*models.Unit, error) {
	return service.repository.FindByID(unitId)
}

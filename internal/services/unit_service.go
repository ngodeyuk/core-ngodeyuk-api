package services

import (
	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/domain/models"
	"ngodeyuk-core/internal/domain/repositories"
)

type UnitService interface {
	Create(dto *dtos.UnitDTO) error
	Update(unitId uint, dto *dtos.UnitDTO) error
	GetAll() ([]models.Unit, error)
	GetByID(unitId uint) (*models.Unit, error)
	DeleteByID(unitId uint) error
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

func (service *unitService) Update(unitId uint, dto *dtos.UnitDTO) error {
	unit, err := service.repository.FindByID(unitId)
	if err != nil {
		return err
	}
	if dto.Title != "" {
		unit.Title = dto.Title
	}

	if dto.Description != "" {
		unit.Description = dto.Description
	}

	if dto.Sequence != 0 {
		unit.Sequence = dto.Sequence
	}
	if err := service.repository.Update(unit); err != nil {
		return err
	}
	return nil
}

func (service *unitService) GetAll() ([]models.Unit, error) {
	return service.repository.FindAll()
}

func (service *unitService) GetByID(unitId uint) (*models.Unit, error) {
	return service.repository.FindByID(unitId)
}

func (service *unitService) DeleteByID(unitId uint) error {
	unit, err := service.repository.FindByID(unitId)
	if err != nil {
		return err
	}
	return service.repository.Delete(unit)
}

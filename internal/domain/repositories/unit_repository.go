package repositories

import (
	"gorm.io/gorm"

	"ngodeyuk-core/internal/domain/models"
)

type UnitRepository interface {
	Create(unit *models.Unit) error
	Update(unit *models.Unit) error
	FindAll() ([]models.Unit, error)
	FindByID(unitId uint) (*models.Unit, error)
	Delete(unit *models.Unit) error
}

type unitRepository struct {
	db *gorm.DB
}

func NewUnitRepository(db *gorm.DB) UnitRepository {
	return &unitRepository{db}
}

func (repository *unitRepository) Create(unit *models.Unit) error {
	return repository.db.Create(unit).Error
}

func (repository *unitRepository) FindAll() ([]models.Unit, error) {
	var units []models.Unit
	if err := repository.db.Find(&units).Error; err != nil {
		return nil, err
	}
	return units, nil
}

func (repository *unitRepository) FindByID(unitId uint) (*models.Unit, error) {
	var unit models.Unit
	if err := repository.db.Where("unit_id = ?", unitId).First(&unit).Error; err != nil {
		return nil, err
	}
	return &unit, nil
}

func (repository *unitRepository) Update(unit *models.Unit) error {
	if err := repository.db.Save(unit).Error; err != nil {
		return err
	}
	return nil
}

func (repository *unitRepository) Delete(unit *models.Unit) error {
	if err := repository.db.Delete(unit).Error; err != nil {
		return err
	}
	return nil
}

package repositories

import (
	"gorm.io/gorm"

	"ngodeyuk-core/internal/domain/models"
)

type UnitRepository interface {
	Create(unit *models.Unit) error
}

type unitRepository struct {
	db *gorm.DB
}

func NewUnitRepository(db *gorm.DB) UnitRepository {
	return &unitRepository{db}
}

func (repository *unitRepository) Create(unit *models.Unit) error {
	return repository.db.Save(unit).Error
}

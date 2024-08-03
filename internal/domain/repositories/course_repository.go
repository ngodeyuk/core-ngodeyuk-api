package repositories

import (
	"gorm.io/gorm"

	"ngodeyuk-core/internal/domain/models"
)

type CourseRepository interface {
	Create(course *models.Course) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

func (repository *courseRepository) Create(course *models.Course) error {
	return repository.db.Create(course).Error
}

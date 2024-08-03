package repositories

import (
	"gorm.io/gorm"

	"ngodeyuk-core/internal/domain/models"
)

// untuk mendefinisikan metode yang diperlukan untuk berinteraksi dengan repository course
type CourseRepository interface {
	// menyimpan data course baru ke database
	Create(course *models.Course) error
}

// untuk mengimplementasikan repository course menggunakan gorm
type courseRepository struct {
	db *gorm.DB
}

// untuk membuat instance baru dengan menginisialisasi koneksi database
func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

// untuk menyimpan data course baru kedalam database menggunakan gorm
func (repository *courseRepository) Create(course *models.Course) error {
	return repository.db.Create(course).Error
}

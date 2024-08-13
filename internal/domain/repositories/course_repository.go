package repositories

import (
	"gorm.io/gorm"

	"ngodeyuk-core/internal/domain/models"
)

// untuk mendefinisikan metode yang diperlukan untuk berinteraksi dengan repository course
type CourseRepository interface {
	// menyimpan data course baru ke database
	Create(course *models.Course) error
	//
	Update(course *models.Course) error
	// mencari semua data course
	FindAll() ([]models.Course, error)
	// mencari data course berdasarkan ID
	FindByID(courseId uint) (*models.Course, error)
	// menghapus data course yang ada didatabase
	Delete(course *models.Course) error
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

// untuk memperbarui data course yang ada didatabase menggunakan gorm
func (repository *courseRepository) Update(course *models.Course) error {
	if err := repository.db.Save(course).Error; err != nil {
		return err
	}
	return nil
}

// untuk mencari semua data course yang ada didatabase menggunakan gorm
func (repository *courseRepository) FindAll() ([]models.Course, error) {
	var courses []models.Course
	if err := repository.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

// untuk mencari data course berdasarkan ID yang ada didatabase menggunakan gorm
func (repository *courseRepository) FindByID(courseId uint) (*models.Course, error) {
	var course models.Course
	if err := repository.db.Where("course_id = ?", courseId).First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

// untuk menghapus data course berdasarkan ID yang ada didatabase menggunakan gorm
func (repository *courseRepository) Delete(course *models.Course) error {
	if err := repository.db.Delete(course).Error; err != nil {
		return err
	}
	return nil
}

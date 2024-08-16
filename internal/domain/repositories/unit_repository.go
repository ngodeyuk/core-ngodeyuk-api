package repositories

import (
	"gorm.io/gorm"

	"ngodeyuk-core/internal/domain/models"
)

// untuk mendefinisikan metode yang diperlukan untuk berinteraksi dengan repository unit
type UnitRepository interface {
	// menyimpan data unit baru ke database
	Create(unit *models.Unit) error
	// untuk memperbarui data unit yg ada didatabase
	Update(unit *models.Unit) error
	// untuk semua data unit yang ada didatabase
	FindAll() ([]models.Unit, error)
	// untuk mencari data unit berdasarkan ID
	FindByID(unitId uint) (*models.Unit, error)
	// untuk menghapus data unit yang ada didatabase
	Delete(unit *models.Unit) error
}

// untuk mengimplementasikan repository unit menggunakan gorm
type unitRepository struct {
	db *gorm.DB
}

// untuk membuat instance baru dengan menginisialisasi koneksi database
func NewUnitRepository(db *gorm.DB) UnitRepository {
	return &unitRepository{db}
}

// untuk menyimpan data unit baru kedalam database menggunakan gorm
func (repository *unitRepository) Create(unit *models.Unit) error {
	return repository.db.Create(unit).Error
}

// untuk mencari semua data unit yang ada didatabase menggunakan gorm
func (repository *unitRepository) FindAll() ([]models.Unit, error) {
	var units []models.Unit
	if err := repository.db.Find(&units).Error; err != nil {
		return nil, err
	}
	return units, nil
}

// untuk mencari data unit berdasarkan id menggunakan gorm
func (repository *unitRepository) FindByID(unitId uint) (*models.Unit, error) {
	var unit models.Unit
	if err := repository.db.Where("unit_id = ?", unitId).First(&unit).Error; err != nil {
		return nil, err
	}
	return &unit, nil
}

// untuk memperbarui data unit yang ada didatabase menggunakan gorm
func (repository *unitRepository) Update(unit *models.Unit) error {
	if err := repository.db.Save(unit).Error; err != nil {
		return err
	}
	return nil
}

// untuk menghapus data unit yang ada didatabase menggunakan gorm
func (repository *unitRepository) Delete(unit *models.Unit) error {
	if err := repository.db.Delete(unit).Error; err != nil {
		return err
	}
	return nil
}

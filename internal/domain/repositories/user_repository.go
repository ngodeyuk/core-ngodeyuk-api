package repositories

import (
	"gorm.io/gorm"

	"ngodeyuk-core/internal/domain/models"
)

// untuk mendefinisikan metode yang diperlukan untuk berinteraksi dengan repository user
type UserRepository interface {
	// menyimpan user baru ke database
	Create(user *models.User) error
	// mencari semua data user
	FindAll() ([]models.User, error)
	// mencari user berdasarkan username
	FindByUsername(username string) (*models.User, error)
	// memperbarui data user yang ada di database
	Update(user *models.User) error
}

// untuk mengimplementasikan repository user mengunakan gorm
type userRepository struct {
	db *gorm.DB
}

// untuk membuat instance baru dengan menginisialisasi koneksi database
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// untuk menyimpan user baru kedalam database menggunakan gorm
func (repository *userRepository) Create(user *models.User) error {
	return repository.db.Create(user).Error
}

// untuk mencari semua data user, lalu mengembalikan semua user jika ditemukan dan error jika tidak ada
func (repository *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := repository.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// untuk mencari user berdasarkan username, lalu mengembalikan user jika ditemukan dan error jika tidak ditemukan
func (repository *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := repository.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// untuk memperbarui data user yang sudah ada didatabase menggunakan gorm
func (repository *userRepository) Update(user *models.User) error {
	if err := repository.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

package repositories

import (
	"gorm.io/gorm"

	"ngodeyuk-core/internal/domain/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	Update(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repository *userRepository) Create(user *models.User) error {
	return repository.db.Create(user).Error
}

func (repository *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := repository.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *userRepository) Update(user *models.User) error {
	if err := repository.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

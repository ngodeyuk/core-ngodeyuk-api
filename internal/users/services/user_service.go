package services

import (
	"errors"
	"ngodeyuk-core/pkg/dto"
	"ngodeyuk-core/pkg/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	RegisterUser(dto.RegisterDTO) (models.User, error)
	LoginUser(dto.LoginDTO) (models.User, error)
	ChangePassword(userID uint, input dto.ChangePasswordDTO) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db}
}

func (s *userService) RegisterUser(registerDTO dto.RegisterDTO) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:     registerDTO.Name,
		Username: registerDTO.Username,
		Password: string(hashedPassword),
	}

	if err := s.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *userService) LoginUser(input dto.LoginDTO) (models.User, error) {
	var user models.User

	username := input.Username
	password := input.Password

	err := s.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}

	if user.UserId == "" {
		return user, errors.New("no user found on that username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) ChangePassword(userID uint, input dto.ChangePasswordDTO) error {
	var user models.User

	err := s.db.Where("user_id = ?", userID).Find(&user).Error
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword))
	if err != nil {
		return errors.New("incorrect old password")
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(newPasswordHash)
	s.db.Save(&user)

	return nil
}

package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/domain/models"
	"ngodeyuk-core/internal/domain/repositories"
	"ngodeyuk-core/pkg/utils"
)

type UserService interface {
	Register(dto *dtos.RegisterDTO) error
	Login(dto *dtos.LoginDTO) (string, error)
	ChangePassword(dto *dtos.ChangePasswordDTO) error
	Update(username string, dto *dtos.UpdateDTO) error
	StartHeartUpdater()
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository}
}

func (service *userService) Register(dto *dtos.RegisterDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(dto.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     dto.Name,
		Username: dto.Username,
		Password: string(hashedPassword),
	}

	return service.repository.Create(user)
}

func (service *userService) Login(dto *dtos.LoginDTO) (string, error) {
	user, err := service.repository.FindByUsername(dto.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = dto.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *userService) ChangePassword(dto *dtos.ChangePasswordDTO) error {
	user, err := service.repository.FindByUsername(dto.Username)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.OldPassword)); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return service.repository.Update(user)
}

func (service *userService) Update(username string, dto *dtos.UpdateDTO) error {
	user, err := service.repository.FindByUsername(username)
	if err != nil {
		return err
	}
	if dto.Name != "" {
		user.Name = dto.Name
	}
	if dto.Point > 0 {
		user.Points += dto.Point
	}
	if dto.Heart > 0 {
		currentHeart := user.Heart
		if currentHeart > 0 {
			user.Heart -= dto.Heart
		}
	}
	user.LastHeartTime = time.Now()
	if err := service.repository.Update(user); err != nil {
		return err
	}
	return nil
}

func (service *userService) StartHeartUpdater() {
	utils.StartHeartUpdater(service.repository, 1*time.Minute)
}

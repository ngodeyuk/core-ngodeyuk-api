package services

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/domain/models"
	"ngodeyuk-core/internal/domain/repositories"
	"ngodeyuk-core/pkg/utils"
)

// mendefinisikan layanan yang diperlukan untuk berinteraksi dengan service
type UserService interface {
	// untuk mendaftarkan pengguna baru berdasarkan dto yang diberikan
	Register(dto *dtos.RegisterDTO) error
	// untuk melakukan proses otentikasi user dan kemudian mengembalikan token JWT apabila berhasil
	Login(dto *dtos.LoginDTO) (string, error)
	// untuk mengubah password user dengan memvalidasi password sebelumnya
	ChangePassword(dto *dtos.ChangePasswordDTO) error
	// untuk mengupdate data user berdasarkan username
	Update(username string, dto *dtos.UpdateDTO) error
	// untuk memulai update otomatis pada heart dengan interval yang sudah ditentukan
	StartHeartUpdater()
	// untuk mengembalikan semua data user
	GetAll() ([]models.User, error)
	// untuk mengembalikan data user berdasarkan username
	GetByUsername(username string) (*models.User, error)
	// untuk menghapus data user berdasarkan username
	DeleteByUsername(username string) error
	// untuk mengupload foto profile pada user berdasarkan username
	UploadProfile(dto *dtos.UploadDTO) error
	// untuk mengembalikan data user berdasarkan point tertinggi
	Leaderboard() ([]models.User, error)
	SelectCourse(username string, courseId uint) error
}

type userService struct {
	repository       repositories.UserRepository
	courseRepository repositories.CourseRepository
}

// untuk membuat instance baru dari user service dengan repository
func NewUserService(
	repository repositories.UserRepository,
	courseRepository repositories.CourseRepository,
) UserService {
	return &userService{repository, courseRepository}
}

func (service *userService) Register(dto *dtos.RegisterDTO) error {
	// validasi ketika username kurang dari 3 karakter
	if len(dto.Username) < 3 {
		return errors.New("username must be atleast 3 character long")
	}
	// validasi ketika password kurang dari 8 karakter
	if len(dto.Password) < 8 {
		return errors.New("password must be atleast 8 character long")
	}

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
	// validasi ketika password lama salah
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.OldPassword)); err != nil {
		return errors.New("old password is incorrect")
	}
	// validasi ketika password lama dengan password baru sama
	if dto.OldPassword == dto.NewPassword {
		return errors.New("new password must be different from old password")
	}
	// validasi ketika password mempunyai kurang dari 8 karakter
	if len(dto.NewPassword) < 8 {
		return errors.New("password must be atleast 8 character long")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return service.repository.Update(user)
}

func (service *userService) Update(username string, dto *dtos.UpdateDTO) error {
	// validasi ketika points memiliki nilai negatif
	if dto.Point < 0 {
		return errors.New("points cant be negative")
	}
	// validasi ketika hearts memiliki nilai negatif
	if dto.Heart < 0 {
		return errors.New("heart cant be negative")
	}
	// validasi ketika pengurangan hearts lebih dari 1
	if dto.Heart > 1 {
		return errors.New("heart cant be reduce more than one")
	}
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

	if dto.Gender != "" {
		user.Gender = dto.Gender
	}

	isMember := user.IsMembership == true
	if !isMember {
		if dto.Heart > 0 {
			currentHeart := user.Heart
			if currentHeart > 0 {
				user.Heart -= dto.Heart
			}
			// validasi ketika heart pada user bernilai 0 maka akan return error
			if currentHeart <= 0 {
				return errors.New("cant reduce heart because current heart value is 0")
			}
		}
	}
	user.LastHeartTime = time.Now()
	if err := service.repository.Update(user); err != nil {
		return err
	}
	return nil
}

func (service *userService) StartHeartUpdater() {
	utils.StartHeartUpdater(service.repository, 1*time.Hour)
}

func (service *userService) GetAll() ([]models.User, error) {
	return service.repository.FindAll()
}

func (service *userService) GetByUsername(username string) (*models.User, error) {
	return service.repository.FindByUsername(username)
}

func (service *userService) DeleteByUsername(username string) error {
	user, err := service.repository.FindByUsername(username)
	if err != nil {
		return err
	}
	return service.repository.Delete(user)
}

func (service *userService) UploadProfile(dto *dtos.UploadDTO) error {
	// menentukan type file pada img yang diupload
	validExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	ext := filepath.Ext(dto.ImgURL)
	// validasi ketika user mengupload file img yang tidak sesuai
	if !validExts[ext] {
		return errors.New("invalid image file type")
	}
	user, err := service.repository.FindByUsername(dto.Username)
	if err != nil {
		return err
	}
	user.ImgURL = dto.ImgURL
	if err := service.repository.Update(user); err != nil {
		return err
	}
	return nil
}

func (service *userService) Leaderboard() ([]models.User, error) {
	users, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].Points > users[j].Points
	})
	var leaderboard []models.User
	for _, user := range users {
		leaderboard = append(leaderboard, models.User{
			Username: user.Username,
			ImgURL:   user.ImgURL,
			Points:   user.Points,
		})
	}
	return leaderboard, nil
}

func (service *userService) SelectCourse(username string, courseId uint) error {
	user, err := service.repository.FindByUsername(username)
	if err != nil {
		return err
	}

	course, err := service.courseRepository.FindByID(courseId)
	if err != nil {
		return err
	}

	user.CourseId = &course.CourseId
	return service.repository.Update(user)
}

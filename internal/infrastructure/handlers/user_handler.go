package handlers

import (
	"fmt"
	"net/http"
	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/services"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByUsername(ctx *gin.Context)
	DeleteByUsername(ctx *gin.Context)
	UploadProfile(ctx *gin.Context)
	Leaderboard(ctx *gin.Context)
	SelectCourse(ctx *gin.Context)
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{service}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dtos.RegisterDTO true "User registration details"
// @Success 201 {object} map[string]interface{}
// @Router /auth/register [post]
func (handler *userHandler) Register(ctx *gin.Context) {
	var input dtos.RegisterDTO
	// validasi ketika requestnya salah/kosong
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	err := handler.service.Register(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"name":     input.Name,
			"username": input.Username,
		},
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dtos.LoginDTO true "User login details"
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
func (handler *userHandler) Login(ctx *gin.Context) {
	var input dtos.LoginDTO
	// validasi ketika requestnya salah/kosong
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	token, err := handler.service.Login(&input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// ChangePassword godoc
// @Summary Change user password
// @Description Update the password for the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param password body dtos.ChangePasswordDTO true "New password details"
// @Success 200 {object} map[string]interface{}
// @Router /user/change-password [put]
func (handler *userHandler) ChangePassword(ctx *gin.Context) {
	var input dtos.ChangePasswordDTO
	// validasi ketika requestnya salah/kosong
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	input.Username = username.(string)
	err := handler.service.ChangePassword(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "change password successfully"})
}

// Update godoc
// @Summary Update user profile
// @Description Update user details such as name, gender, points, and heart
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param user body dtos.UpdateDTO true "Updated user details"
// @Success 200 {object} map[string]interface{}
// @Router /user/update [patch]
func (handler *userHandler) Update(ctx *gin.Context) {
	var input dtos.UpdateDTO
	// validasi ketika requesnya salah/kosong
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	// validasi ketika request tidak diisi sama sekali/kosong
	if input.Name == "" && input.Gender == "" && input.Point == 0 && input.Heart == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "at least one field must be provided"})
		return
	}
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := handler.service.Update(username.(string), &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"name":   input.Name,
			"point":  input.Point,
			"heart":  input.Heart,
			"gender": input.Gender,
		},
	})
}

// GetAll godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags user
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} dtos.UserDTO
// @Router /user [get]
func (handler *userHandler) GetAll(ctx *gin.Context) {
	users, err := handler.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response []dtos.UserDTO
	for _, user := range users {
		response = append(response, dtos.UserDTO{
			UserId:       user.UserId,
			Name:         user.Name,
			Username:     user.Username,
			ImgURL:       "/" + user.ImgURL,
			Heart:        user.Heart,
			Points:       user.Points,
			Gender:       user.Gender,
			IsMembership: user.IsMembership,
			IsAdmin:      user.IsAdmin,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// GetByUsername godoc
// @Summary Get user by username
// @Description Retrieve details of a user by their username
// @Tags user
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} dtos.UserDTO
// @Router /user/current [get]
func (handler *userHandler) GetByUsername(ctx *gin.Context) {
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	user, err := handler.service.GetByUsername(username.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	response := dtos.UserDTO{
		UserId:       user.UserId,
		Name:         user.Name,
		ImgURL:       "/" + user.ImgURL,
		Username:     user.Username,
		Heart:        user.Heart,
		Points:       user.Points,
		Gender:       user.Gender,
		IsMembership: user.IsMembership,
		IsAdmin:      user.IsAdmin,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// DeleteByUsername godoc
// @Summary Delete user by username
// @Description Delete a user account based on the username
// @Tags user
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} map[string]interface{}
// @Router /user/delete [delete]
func (handler *userHandler) DeleteByUsername(ctx *gin.Context) {
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	err := handler.service.DeleteByUsername(username.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "delete user successfully"})
}

// UploadProfile godoc
// @Summary Upload user profile image
// @Description Upload a new profile image for the authenticated user
// @Tags user
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param image formData file true "Profile image file"
// @Success 200 {object} map[string]interface{}
// @Router /user/upload [post]
func (handler *userHandler) UploadProfile(ctx *gin.Context) {
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "image file required"})
		return
	}
	fileExt := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", username.(string), fileExt)
	filepath := "public/users/" + filename

	if _, err := os.Stat(filepath); err != nil {
		if err := os.Remove(filepath); err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "failed to delete exiting images"},
			)
		}
	}

	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	input := dtos.UploadDTO{
		Username: username.(string),
		ImgURL:   filepath,
	}
	if err := handler.service.UploadProfile(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// Leaderboard godoc
// @Summary Get leaderboard
// @Description Retrieve the leaderboard with user points and ranking
// @Tags user
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} dtos.LeaderboardDTO
// @Router /user/leaderboard [get]
func (handler *userHandler) Leaderboard(ctx *gin.Context) {
	leaderboard, err := handler.service.Leaderboard()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response []dtos.LeaderboardDTO
	for _, user := range leaderboard {
		response = append(response, dtos.LeaderboardDTO{
			Username: user.Username,
			ImgURL:   "/" + user.ImgURL,
			Points:   user.Points,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// SelectCourse godoc
// @Summary Select a course for the user
// @Description Allows the authenticated user to select a course by ID
// @Tags user
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param course_id path integer true "Course ID"
// @Success 200 {object} map[string]interface{}
// @Router /user/select-course/{course_id} [post]
func (handler *userHandler) SelectCourse(ctx *gin.Context) {
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	courseIdstr := ctx.Param("course_id")
	courseId, err := strconv.ParseUint(courseIdstr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	err = handler.service.SelectCourse(username.(string), uint(courseId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user successfully to select course"})
}

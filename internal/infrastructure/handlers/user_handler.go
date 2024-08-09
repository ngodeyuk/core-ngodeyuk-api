package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/services"
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
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{service}
}

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

func (handler *userHandler) Update(ctx *gin.Context) {
	var input dtos.UpdateDTO
	// validasi ketika requesnya salah/kosong
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	// validasi ketika request tidak diisi sama sekali/kosong
	if input.Name == "" && input.Point == 0 && input.Heart == 0 {
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
			"name":  input.Name,
			"point": input.Point,
			"heart": input.Heart,
		},
	})
}

func (handler *userHandler) GetAll(ctx *gin.Context) {
	users, err := handler.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response []dtos.UserDTO
	for _, user := range users {
		response = append(response, dtos.UserDTO{
			UserId:   user.UserId,
			Name:     user.Name,
			Username: user.Username,
			ImgURL:   "/" + user.ImgURL,
			Heart:    user.Heart,
			Points:   user.Points,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

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
		UserId:   user.UserId,
		Name:     user.Name,
		ImgURL:   "/" + user.ImgURL,
		Username: user.Username,
		Heart:    user.Heart,
		Points:   user.Points,
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

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

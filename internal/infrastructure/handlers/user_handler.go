package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/services"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{service}
}

func (handler *userHandler) Register(ctx *gin.Context) {
	var input dtos.RegisterDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

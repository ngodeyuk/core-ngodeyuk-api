package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/services"
)

type CourseHandler interface {
	Create(ctx *gin.Context)
}

type courseHandler struct {
	service services.CourseService
}

func NewCourseHandler(service services.CourseService) CourseHandler {
	return &courseHandler{service}
}

func (handler courseHandler) Create(ctx *gin.Context) {
	var input dtos.CourseDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := handler.service.Create(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"name": input.Title,
			"img":  input.Img,
		},
	})
}

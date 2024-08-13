package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/services"
)

type UnitHandler interface {
	Create(ctx *gin.Context)
}

type unitHandler struct {
	service services.UnitService
}

func NewUnitHandler(service services.UnitService) UnitHandler {
	return &unitHandler{service}
}

func (handler *unitHandler) Create(ctx *gin.Context) {
	var input dtos.UnitDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.service.Create(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"title":       input.Title,
			"description": input.Description,
			"sequence":    input.Sequence,
		},
	})
}

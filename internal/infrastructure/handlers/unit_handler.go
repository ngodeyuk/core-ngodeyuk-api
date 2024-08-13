package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/services"
)

type UnitHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
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

func (handler *unitHandler) Update(ctx *gin.Context) {
	var input dtos.UnitDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request payload"})
		return
	}
	unitIdStr := ctx.Param("unit_id")
	unitId, err := strconv.ParseUint(unitIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid unit ID"})
		return
	}
	unit, err := handler.service.GetByID(uint(unitId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	unit.Title = input.Title
	unit.Description = input.Description
	unit.Sequence = input.Sequence

	err = handler.service.Update(uint(unitId), &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "update unit successfuly"})
}

func (handler *unitHandler) GetAll(ctx *gin.Context) {
	units, err := handler.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dtos.UnitDTO
	for _, unit := range units {
		response = append(response, dtos.UnitDTO{
			UnitId:      unit.UnitId,
			Title:       unit.Title,
			Description: unit.Description,
			Sequence:    unit.Sequence,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (handler *unitHandler) GetByID(ctx *gin.Context) {
	unitIdStr := ctx.Param("unit_id")
	unitId, err := strconv.ParseUint(unitIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid unit ID"})
		return
	}

	unit, err := handler.service.GetByID(uint(unitId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dtos.UnitDTO{
		UnitId:      unit.UnitId,
		Title:       unit.Title,
		Description: unit.Description,
		Sequence:    unit.Sequence,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (handler *unitHandler) DeleteByID(ctx *gin.Context) {
	unitIdStr := ctx.Param("unit_id")
	unitId, err := strconv.ParseUint(unitIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid unit ID"})
		return
	}
	err = handler.service.DeleteByID(uint(unitId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "delete unit successfuly"})
}

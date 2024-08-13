package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/services"
)

type CourseHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type courseHandler struct {
	service services.CourseService
}

func NewCourseHandler(service services.CourseService) CourseHandler {
	return &courseHandler{service}
}

func (handler *courseHandler) Create(ctx *gin.Context) {
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

func (handler *courseHandler) Update(ctx *gin.Context) {
	var input dtos.CourseDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request payload"})
		return
	}
	courseIdstr := ctx.Param("course_id")
	courseId, err := strconv.ParseUint(courseIdstr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}
	course, err := handler.service.GetByID(uint(courseId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	course.Title = input.Title
	course.Img = input.Img

	err = handler.service.Update(uint(courseId), &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "update course successfully"})
}

func (handler *courseHandler) GetAll(ctx *gin.Context) {
	courses, err := handler.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dtos.CourseDTO
	for _, course := range courses {
		response = append(response, dtos.CourseDTO{
			CourseId: course.CourseId,
			Title:    course.Title,
			Img:      course.Img,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (handler *courseHandler) GetByID(ctx *gin.Context) {
	courseIdstr := ctx.Param("course_id")
	courseId, err := strconv.ParseUint(courseIdstr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}
	course, err := handler.service.GetByID(uint(courseId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := dtos.CourseDTO{
		CourseId: course.CourseId,
		Title:    course.Title,
		Img:      course.Img,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (handler *courseHandler) DeleteByID(ctx *gin.Context) {
	courseIdstr := ctx.Param("course_id")
	courseId, err := strconv.ParseUint(courseIdstr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	err = handler.service.DeleteByID(uint(courseId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "delete course successfully"})
}

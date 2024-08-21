package handlers

import (
	"net/http"
	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
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

// Create Course godoc
// @Summary Create a new course
// @Description Create a new course with the provided details.
// @Tags Course
// @Accept json
// @Produce json
// @Param course body dtos.CourseDTO true "Create course data details"
// @Success 201 {object} map[string]interface{}
// @Router /api/course [post]
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
	ctx.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"name": input.Title,
			"img":  input.Img,
		},
	})
}

// Update Course godoc
// @Summary Update an existing course
// @Description Update an existing course identified by course ID with the provided details.
// @Tags Course
// @Accept json
// @Produce json
// @Param course_id path int true "Course ID"
// @Param course body dtos.CourseDTO true "Updated course data details"
// @Success 200 {object} map[string]interface{}
// @Router /api/course/{course_id} [patch]
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

// GetAll Courses godoc
// @Summary Retrieve all courses
// @Description Retrieve a list of all courses.
// @Tags Course
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/course [get]
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

// GetByID Course godoc
// @Summary Retrieve a course by ID
// @Description Retrieve a course identified by course ID.
// @Tags Course
// @Produce json
// @Param course_id path int true "Course ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/course/{course_id} [get]
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

// DeleteByID Course godoc
// @Summary Delete a course by ID
// @Description Delete a course identified by course ID.
// @Tags Course
// @Param course_id path int true "Course ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/course/{course_id} [delete]
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

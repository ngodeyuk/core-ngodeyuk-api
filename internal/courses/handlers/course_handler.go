package handlers

import (
	"net/http"
	"ngodeyuk-core/internal/courses/services"
	"ngodeyuk-core/pkg/dto"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService services.CourseService
}

func NewCourseHandler(courseService services.CourseService) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var courseDTO dto.CreateCourseDTO
	if err := c.ShouldBindJSON(&courseDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.courseService.CreateCourse(courseDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

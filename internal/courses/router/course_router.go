package router

import (
	"ngodeyuk-core/internal/courses/handlers"
	"ngodeyuk-core/internal/courses/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CourseRouter(r *gin.Engine, db *gorm.DB) {
	couseService := services.NewCourseService(db)
	userHandler := handlers.NewCourseHandler(couseService)

	r.POST("api/course/create", userHandler.CreateCourse)
}

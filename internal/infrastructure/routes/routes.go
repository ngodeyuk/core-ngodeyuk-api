package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ngodeyuk-core/internal/domain/repositories"
	"ngodeyuk-core/internal/infrastructure/handlers"
	"ngodeyuk-core/internal/services"
)

func SetupRoutes(route *gin.Engine, db *gorm.DB) {
	repository := repositories.NewUserRepository(db)
	service := services.NewUserService(repository)
	handler := handlers.NewUserHandler(service)

	route.POST("auth/register", handler.Register)
	route.POST("auth/login", handler.Login)

	api := route.Group("api")
	{
		repository := repositories.NewCourseRepository(db)
		service := services.NewCourseService(repository)
		handler := handlers.NewCourseHandler(service)

		api.POST("course", handler.Create)

	}
}

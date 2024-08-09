package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ngodeyuk-core/internal/domain/repositories"
	"ngodeyuk-core/internal/infrastructure/handlers"
	"ngodeyuk-core/internal/infrastructure/middleware"
	"ngodeyuk-core/internal/services"
)

func SetupRoutes(route *gin.Engine, db *gorm.DB) {
	repository := repositories.NewUserRepository(db)
	service := services.NewUserService(repository)
	service.StartHeartUpdater()
	handler := handlers.NewUserHandler(service)

	route.POST("auth/register", handler.Register)
	route.POST("auth/login", handler.Login)

	user := route.Group("user")
	user.Use(middleware.AuthMiddleware())
	{
		user.PUT("upload", handler.UploadProfile)
		user.PUT("change-password", handler.ChangePassword)
		user.PATCH("update", handler.Update)
		user.GET("/", handler.GetAll)
		user.GET("current", handler.GetByUsername)
		user.GET("leaderboard", handler.Leaderboard)
		user.DELETE("delete", handler.DeleteByUsername)
	}

	api := route.Group("api")
	{
		repository := repositories.NewCourseRepository(db)
		service := services.NewCourseService(repository)
		handler := handlers.NewCourseHandler(service)

		api.POST("course", handler.Create)

	}
}

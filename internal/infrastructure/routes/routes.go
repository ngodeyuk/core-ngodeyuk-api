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
	userRepository := repositories.NewUserRepository(db)
	courseRepository := repositories.NewCourseRepository(db)
	unitRepository := repositories.NewUnitRepository(db)

	userService := services.NewUserService(userRepository, courseRepository)
	userService.StartHeartUpdater()
	courseService := services.NewCourseService(courseRepository)
	unitService := services.NewUnitService(unitRepository)

	userHandler := handlers.NewUserHandler(userService)
	courseHandler := handlers.NewCourseHandler(courseService)
	UnitHandler := handlers.NewUnitHandler(unitService)

	route.POST("auth/register", userHandler.Register)
	route.POST("auth/login", userHandler.Login)

	user := route.Group("user")
	user.Use(middleware.AuthMiddleware())
	{
		user.PUT("upload", userHandler.UploadProfile)
		user.PUT("change-password", userHandler.ChangePassword)
		user.PATCH("update", userHandler.Update)
		user.GET("/", userHandler.GetAll)
		user.GET("current", userHandler.GetByUsername)
		user.GET("leaderboard", userHandler.Leaderboard)
		user.POST("select-course/:course_id", userHandler.SelectCourse)
		user.DELETE("delete", userHandler.DeleteByUsername)
	}

	api := route.Group("api")
	{
		api.POST("course", courseHandler.Create)
		api.PATCH("course/:course_id", courseHandler.Update)
		api.GET("course", courseHandler.GetAll)
		api.GET("course/:course_id", courseHandler.GetByID)
		api.DELETE("course/:course_id", courseHandler.DeleteByID)

		api.POST("unit", UnitHandler.Create)
		api.PATCH("unit/:unit_id", UnitHandler.Update)
		api.GET("unit", UnitHandler.GetAll)
		api.GET("unit/:unit_id", UnitHandler.GetByID)
		api.DELETE("unit/:unit_id", UnitHandler.DeleteByID)
	}
}

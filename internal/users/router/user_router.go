package router

import (
	"ngodeyuk-core/internal/users/handlers"
	"ngodeyuk-core/internal/users/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(r *gin.Engine, db *gorm.DB) {
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	r.POST("api/register", userHandler.RegisterUser)
}

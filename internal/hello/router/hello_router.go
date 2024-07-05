package router

import (
	"ngodeyuk-core/internal/hello/handlers"

	"github.com/gin-gonic/gin"
)

func HelloRouter(r *gin.Engine) {
	r.GET("api/hello", handlers.HelloHandler)
}

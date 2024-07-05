package router

import (
	"ngodeyuk-core/internal/hello/handlers"

	"github.com/gin-gonic/gin"
)

func HelloRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", handlers.HelloHandler)
	return r
}

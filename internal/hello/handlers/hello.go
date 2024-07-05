package handlers

import (
	"net/http"
	"ngodeyuk-core/internal/hello/services"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	message := services.GetHelloMessage()
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

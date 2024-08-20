package main

import (
	"fmt"
	"log"
	_ "ngodeyuk-core/cmd/docs"
	"ngodeyuk-core/database"
	"ngodeyuk-core/internal/infrastructure/routes"
	"ngodeyuk-core/pkg/utils"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	utils.LoadEnv()
}

func main() {
	route := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	route.Use(cors.New(config))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database.")
	}

	publicPath := filepath.Join(".", "public")
	route.Static("/public", publicPath)

	routes.SetupRoutes(route, db)

	fmt.Println("Server is running at http://localhost:2000")
	log.Fatal(route.Run(":2000"))
}

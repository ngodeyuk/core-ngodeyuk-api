package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ngodeyuk-core/database"
	"ngodeyuk-core/internal/infrastructure/routes"
)

func main() {
	route := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}

	route.Use(cors.New(config))

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error load .env file: %v", err)
	}

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

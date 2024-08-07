package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ngodeyuk-core/database"
	"ngodeyuk-core/internal/infrastructure/routes"
)

func main() {
	route := gin.Default()
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

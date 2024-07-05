package main

import (
	"fmt"
	"log"
	"net/http"
	"ngodeyuk-core/config"
	"ngodeyuk-core/internal/hello/router"
	"ngodeyuk-core/migrations"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	err = migrations.Migrate(db)
	if err != nil {
		log.Fatal("Failed to migrate database schema")
	}

	r := router.HelloRouter()

	fmt.Println("Server is running at http://localhost:2000")
	log.Fatal(http.ListenAndServe(":2000", r))
}

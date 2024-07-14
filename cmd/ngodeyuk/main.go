package main

import (
	"fmt"
	"log"
	"ngodeyuk-core/config"
	courseRouter "ngodeyuk-core/internal/courses/router"
	helloRouter "ngodeyuk-core/internal/hello/router"
	userRouter "ngodeyuk-core/internal/users/router"
	"ngodeyuk-core/migrations"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	helloRouter.HelloRouter(r)
	userRouter.UserRouter(r, db)
	courseRouter.CourseRouter(r, db)

	fmt.Println("Server is running at http://localhost:2000")
	log.Fatal(r.Run(":2000"))
}

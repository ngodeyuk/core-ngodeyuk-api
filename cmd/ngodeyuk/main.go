package main

import (
	"fmt"
	"log"
	"net/http"
	"ngodeyuk-core/internal/hello/router"
)

func main() {
	r := router.HelloRouter()

	fmt.Println("Server is running at http://localhost:2000")
	log.Fatal(http.ListenAndServe(":2000", r))
}

package main

import (
	"log"
	"net/http"
	"server/internal/routes"
)

func main() {
	router := routes.SetupRoutes()
	log.Println("Server is running on port: 8080")
	http.ListenAndServe(":8080", router)
}

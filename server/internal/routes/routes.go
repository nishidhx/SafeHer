package routes

import (
	"server/internal/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	authHandler := handlers.NewAuthHandler(db)
	// Register the auth handler for the "/auth" route (example)
	router.HandleFunc("/auth", authHandler.SafeHerLogin).Methods("POST")

	return router
}

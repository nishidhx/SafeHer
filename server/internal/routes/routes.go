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
	router.HandleFunc("/auth/login", authHandler.SafeHerLogin).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/register", authHandler.SafeHerRegister).Methods("POST", "OPTIONS")

	return router
}

package main

import (
	"log"
	"net/http"
	"server/internal/database"
	"server/internal/models"
	"server/internal/routes"
)

func main() {
	db := database.Connect()

	if db != nil {
		log.Println("DB connected")
	}

	/* Auto migration of tables */
	db.AutoMigrate(
		&models.User{},
		&models.EmergencyContact{},
		&models.RiskZone{},
		&models.SOSAlert{},
		&models.Incident{},
	)

	router := routes.SetupRoutes(db)
	log.Println("Server is running on port: 8080")
	http.ListenAndServe(":8080", router)
}

package handlers

import "gorm.io/gorm"

type IncidentHandler struct {
	DB *gorm.DB
}

func NewIncidentHandler(db *gorm.DB) *IncidentHandler {
	return &IncidentHandler{
		DB: db,
	}
}

func _()

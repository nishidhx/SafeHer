package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/models"
	"server/pkg/conversions"

	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		DB: db,
	}
}

func (h *AuthHandler) SafeHerLogin(writer http.ResponseWriter, request *http.Request) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// decodes the json from the request body and returns the err if error comes.
	err := json.NewDecoder(request.Body).Decode(&req)

	if err != nil {
		http.Error(writer, "Invalid Request body", http.StatusBadRequest)
		log.Println("requst body not found or might be invalid")
		return
	}

	var user models.User

	result := h.DB.Where("email = ?", req.Email).First(&user)

	if result.Error != nil {
		http.Error(writer, "user not found", http.StatusUnauthorized)
		log.Fatalln("User not found")
		return
	}

	if user.Password != req.Password {
		http.Error(writer, "password or email is incorrect", http.StatusUnauthorized)
		log.Fatalln("password is incorrect")
	}

	json.NewEncoder(writer).Encode(user)
}

func (h *AuthHandler) SafeHerRegister(writer http.ResponseWriter, request *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// decode the json from the request body for the signup endpoint and returns the err if error comes.
	err := json.NewDecoder(request.Body).Decode(&req)

	if err != nil {
		http.Error(writer, "Invalid Request body", http.StatusBadRequest)
		log.Println("requst body not found or might be invalid")
		return
	}

	var hashed_pass string = conversions.HashingService(req.Password)

	user := models.User{
		Email:    req.Email,
		Password: hashed_pass,
	}

	result := h.DB.Create(&user)

	if result.Error != nil {
		http.Error(writer, "failed to create user", http.StatusUnauthorized)
		log.Println("Internal server error from db")
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("User registered successfully"))
}

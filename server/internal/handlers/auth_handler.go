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

	isPasswordMatched := conversions.ComparehashPassword(req.Password, user.Password)

	if isPasswordMatched == false {
		http.Error(writer, "password or email is incorrect", http.StatusUnauthorized)
		log.Println("password is incorrect")
		return
	}

	json.NewEncoder(writer).Encode(map[string]string{
		"message": "Login successful",
		"id":      string(user.ID),
	})
}

func (h *AuthHandler) SafeHerRegister(writer http.ResponseWriter, request *http.Request) {
	var req struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		PhoneNumber string `json:"phone_number"`
		Name        string `json:"name"`
	}

	// decode the json from the request body for the signup endpoint and returns the err if error comes.
	err := json.NewDecoder(request.Body).Decode(&req)

	if err != nil {
		http.Error(writer, "Invalid Request body", http.StatusBadRequest)
		log.Println("requst body not found or might be invalid")
		return
	}

	log.Printf(req.Email, req.Password)

	var hashed_pass string = conversions.HashingService(req.Password)

	user := models.User{
		Email:       req.Email,
		Password:    hashed_pass,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	log.Printf("%s", user.Email)

	result := h.DB.Create(&user)

	if result.Error != nil {
		http.Error(writer, "failed to create user", http.StatusUnauthorized)
		log.Println("Internal server error from db: ", result)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("User registered successfully"))

}

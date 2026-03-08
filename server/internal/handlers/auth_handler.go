package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/models"
	"server/pkg/conversions"
	"server/pkg/jwt"
	"time"

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

	tokenService, err := jwt.NewTokenService("SAFE_HER_SECRET")

	emergencyContactStrings := make([]string, len(user.EmergencyContacts))
	for i, contact := range user.EmergencyContacts {
		emergencyContactStrings[i] = contact.PhoneNumber // or appropriate field
	}

	userAuthenticationToken, err := tokenService.GenerateUserAuthenticationToken(jwt.TokenPayload{
		UserID:           user.ID,
		Email:            req.Email,
		Username:         user.Name,
		EmergencyContact: emergencyContactStrings,
	})

	/* Cookie Logic */
	expiration := time.Now().Add(24 * time.Hour) // Cookie expires in 24 hours
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    userAuthenticationToken,
		Expires:  expiration,
		Path:     "/",                  // Makes the cookie available across the entire site
		HttpOnly: true,                 // Prevents access via JavaScript
		Secure:   false,                // Ensures the cookie is sent over HTTPS only (recommended for production)
		SameSite: http.SameSiteLaxMode, // Provides protection against CSRF
	}

	http.SetCookie(writer, &cookie)
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(map[string](string){
		"message": "Login successful",
		"id":      string(user.ID),
		"token":   userAuthenticationToken,
	})
	writer.Write([]byte("Cookie has been set!"))
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

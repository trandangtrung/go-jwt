// controllers/auth_controller.go
package controllers

import (
	"JWT/repositories"
	"JWT/services"
	"encoding/json"
	"net/http"
)

var userRepository = repositories.NewUserRepository()

func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := userRepository.FindByUsername(creds.Username)
	if err != nil || user.Password != creds.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := services.CreateToken(user.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Login successful for user: " + creds.Username))
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

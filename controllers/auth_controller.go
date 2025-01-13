// // controllers/auth_controller.go
// package controllers

// import (
// 	"JWT/repositories"
// 	"JWT/services"
// 	"encoding/json"
// 	"net/http"
// )

// var userRepository = repositories.NewUserRepository()

// var users = make(map[string]string) // đơn giản hóa với map để lưu trữ người dùng

// func Register(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var user User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	if _, exists := users[user.Username]; exists {
// 		http.Error(w, "User already exists", http.StatusConflict)
// 		return
// 	}

// 	users[user.Username] = user.Password
// 	w.Write([]byte("Registration successful"))
// }

// func Login(w http.ResponseWriter, r *http.Request) {
// 	var creds struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := userRepository.FindByUsername(creds.Username)
// 	if err != nil || user.Password != creds.Password {
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	token, err := services.CreateToken(user.ID)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write([]byte("Login successful for user: " + creds.Username))
// 	json.NewEncoder(w).Encode(map[string]string{"token": token})
// }

package controllers

import (
	"JWT/repositories"
	"JWT/services"
	"encoding/json"
	"net/http"
)

var userRepo = repositories.NewUserRepository()

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := userRepo.AddUser(user.Username, user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Write([]byte("Registration successful"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if userRepo.ValidateCredentials(user.Username, user.Password) {
		token, err := services.CreateToken(user.Username)
		if err != nil {
			http.Error(w, "Could not create token", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Login successful. Token: " + token))
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Xử lý logout (có thể đơn giản là xóa token phía client)
	w.Write([]byte("Logout successful"))
}

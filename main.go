// main.go
package main

import (
	"JWT/controllers"
	"JWT/middlewares"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", controllers.Login)

	// Protected route
	http.Handle("/protected", middlewares.AuthenticateToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a protected route"))
	})))

	// Khởi động server
	log.Println("Server is starting on port 3050...")
	if err := http.ListenAndServe(":3050", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

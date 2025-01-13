//

package main

import (
	"JWT/controllers"
	middleware "JWT/middleware"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)

	// Middleware cho các route cần xác thực
	http.HandleFunc("/protected", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value("username").(string)
		w.Write([]byte("Protected route accessed by: " + username))
	}))

	log.Println("Server is starting on port 3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

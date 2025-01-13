//

package middleware

import (
	"JWT/services"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := services.VerifyToken(strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 	// Lưu thông tin người dùng vào context nếu cần
		ctx := context.WithValue(r.Context(), "username", claims["username"])
		r = r.WithContext(ctx) // Cập nhật request với context mới
		next.ServeHTTP(w, r)
	}
}

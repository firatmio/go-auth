package main

import (
	"log"
	"net/http"
	"strings"

	"auth/handlers"
	"auth/models"
	"auth/utils"
)

func main() {
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/home", authMiddleware(handlers.Home))
	http.HandleFunc("/users", authMiddleware(requirePermission(models.PermRead, handlers.ListUsers)))
	http.HandleFunc("/admin", authMiddleware(requirePermission(models.PermAdmin, handlers.Home)))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requirePermission(requiredPerm models.Permission, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)
		claims, _ := utils.ValidateToken(tokenString) // Error checked in authMiddleware already

		// Check if user has the required permission bit set
		if (claims.Permissions & requiredPerm) != requiredPerm {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden: Insufficient permissions"))
			return
		}

		next(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	cookie, err := r.Cookie("token")
	if err == nil {
		return cookie.Value
	}
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			return bearerToken[1]
		}
	}
	return ""
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token from cookie or header
		tokenString := getTokenFromRequest(r)

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// You can add claims to context here if needed
		log.Printf("User %s (Perms: %d) accessed %s", claims.Username, claims.Permissions, r.URL.Path)

		next(w, r)
	}
}

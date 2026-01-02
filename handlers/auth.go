package handlers

import (
	"encoding/json"
	"net/http"

	"auth/models"
	"auth/utils"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username    string            `json:"username"`
	Password    string            `json:"password"`
	Permissions models.Permission `json:"permissions"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = models.CreateUser(creds.Username, string(hashedPassword), creds.Permissions)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := models.GetUser(creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateToken(creds.Username, user.Permissions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: utils.GetExpirationTime(), // I need to expose this or just use a fixed time here
	})

	// Also return it in body for convenience
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the protected home page!"))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

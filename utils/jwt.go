package utils

import (
	"auth/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username    string            `json:"username"`
	Permissions models.Permission `json:"permissions"`
	jwt.RegisteredClaims
}

func GetExpirationTime() time.Time {
	return time.Now().Add(5 * time.Minute)
}

func GenerateToken(username string, permissions models.Permission) (string, error) {
	expirationTime := GetExpirationTime()
	claims := &Claims{
		Username:    username,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

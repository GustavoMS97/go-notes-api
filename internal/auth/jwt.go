package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID string, isRefresh bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
	}

	if isRefresh {
		claims["refresh"] = true
		claims["exp"] = time.Now().Add(7 * 24 * time.Hour).Unix() // 7 dias
	} else {
		claims["exp"] = time.Now().Add(30 * time.Minute).Unix() // 30 minutos
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

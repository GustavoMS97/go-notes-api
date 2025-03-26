package auth

import (
	"errors"
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
		claims["exp"] = time.Now().Add(7 * 24 * time.Hour).Unix()
	} else {
		claims["exp"] = time.Now().Add(30 * time.Minute).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseAndValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func ParseAndValidateRefreshToken(tokenStr string) (string, error) {
	claims, err := ParseAndValidateJWT(tokenStr)
	if err != nil {
		return "", err
	}

	refresh, ok := claims["refresh"].(bool)
	if !ok || !refresh {
		return "", errors.New("not a refresh token")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid user id in token")
	}

	return userID, nil
}

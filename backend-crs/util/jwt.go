package util

import (
	"backend-crs/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(config.GetSecretKey())

func ValidateJwtToken(jwtToken string) (map[string]string, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	userId, ok1 := claims["userId"].(string)
	if !ok1 {
		return nil, errors.New("userId claim not found or not a string")
	}

	role, ok2 := claims["role"].(string)
	if !ok2 {
		return nil, errors.New("role claim not found or not a string")
	}

	return map[string]string{
		"userId": userId,
		"role":   role,
	}, nil
}

func GenerateJwtToken(userId string, role string) (string, error) {

	claims := jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

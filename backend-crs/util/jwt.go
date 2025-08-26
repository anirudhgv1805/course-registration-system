package util

import (
	"backend-crs/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(config.GetSecretKey())


func ValidateJwtToken(jwtToken string) (string, error) {

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error ) {
		return token, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		registerNo, ok := claims["register_no"].(string)
		if !ok {
			return "",jwt.ErrInvalidKey
		}
		return registerNo, nil
	}
	
	return "", err
}

func GenerateJwtToken(registerNo string, role string) (string, error) {

	claims := jwt.MapClaims{
		"register_no": registerNo ,
		"exp" : time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

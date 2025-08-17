package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(hpass) ,err
}
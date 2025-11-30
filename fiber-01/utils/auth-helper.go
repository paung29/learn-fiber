package utils

import (
	"golang.org/x/crypto/bcrypt"
)


func HashedPassword(password string) (string, error) {
	code, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return  string(code), err
}
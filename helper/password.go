package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(hash, password string) (result bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return
	}
	return true, nil
}

func HashPassword(password string) (result string, err error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	result = string(byte)
	return
}

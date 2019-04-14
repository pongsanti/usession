package usession

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordByteArray := toByteArray(password)
	hashByteArray, err := bcrypt.GenerateFromPassword(passwordByteArray, bcrypt.DefaultCost)
	if err != nil {
		log.Print("Error create hash ", err)
		return "", err
	}
	return toString(hashByteArray), nil
}

func ComparePasswordWithHash(password string, hash string) bool {
	passwordByteArray := toByteArray(password)
	hashByteArray := toByteArray(hash)

	err := bcrypt.CompareHashAndPassword(hashByteArray, passwordByteArray)
	if err != nil {
		log.Print("Error compare hash ", err)
		return false
	}
	return true
}

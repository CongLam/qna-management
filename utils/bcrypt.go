package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pw := []byte(password)
	res, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		GetLogger().Fatal(err.Error())
	}
	return string(res)
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hash := []byte(hashPassword)
	return bcrypt.CompareHashAndPassword(hash, pw)
}

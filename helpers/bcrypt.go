package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		log.Fatal("Error comparing: ", err)
	}

	return err == nil
}
package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(hashedPassword string, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}

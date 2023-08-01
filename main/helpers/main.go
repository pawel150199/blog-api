package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func CalculatePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

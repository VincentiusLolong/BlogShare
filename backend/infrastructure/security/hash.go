package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	hashedPassword := string(salt)
	return hashedPassword, nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

package utils

import "golang.org/x/crypto/bcrypt"

// Hashear un string, por defecto son 10 salts
func HashString(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 10)
	return string(bytes), err
}

// Verificacion normal de un hash y un input
func VerifyHash(input, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
	return err == nil
}

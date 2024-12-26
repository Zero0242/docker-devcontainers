package utils

import "net/mail"

// Validamos si el email proveido es valido o no
func ValidateEmail(input string) bool {
	_, err := mail.ParseAddress(input)
	return err != nil
}

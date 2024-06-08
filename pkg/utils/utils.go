package utils

import (
	"net/mail"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func IsMailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

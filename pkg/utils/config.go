package utils

import (
	"os"
)

const (
	adminEmail    = "aaaaaaaa@aaaaaaaa.ch"
	connString    = "host=localhost port=5432 user=banague password=banaguePassword! dbname=baby_names sslmode=disable"
	eventPassword = "mantis-shrimp"
	twint         = "0796661337"
	paypal        = "mail@domain.com"
	sessionKey    = "secret-session-key-1337!"
)

// GetAdminEmail returns the email address of the admin.
func GetAdminEmail() string {
	return GetEnv("ADMIN_EMAIL", adminEmail)
}

// GetDatabaseConnectionString returns the connection string for the database.
func GetDatabaseConnectionString() string {
	return GetEnv("DATABASE_URL", connString)
}

// GetEventPassword returns the password for the event.
func GetEventPassword() string {
	return GetEnv("EVENT_PASSWORD", eventPassword)
}

func GetMobile() string {
	return GetEnv("TWINT", twint)
}

func GetPaypal() string {
	return GetEnv("PAYPAL", paypal)
}

func GetSessionKey() string {
	return GetEnv("SESSION_KEY", sessionKey)
}

// GetEnv returns the value of an environment variable or a fallback value if the variable is not set.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

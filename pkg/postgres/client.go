package postgres

import (
	"database/sql"
	"fmt"

	"sbebe.ch/baby-name-guesser/pkg/utils"
)

const (
	connString = "host=localhost port=5432 user=banague password=banaguePassword! dbname=baby_names sslmode=disable"
)

type PostgresClient struct {
	db *sql.DB
}

func NewPostgresClient() (*PostgresClient, error) {
	connString := utils.GetEnv("DATABASE_URL", connString)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// err = db.Ping()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to ping the database: %v", err)
	// }

	return &PostgresClient{db: db}, nil
}

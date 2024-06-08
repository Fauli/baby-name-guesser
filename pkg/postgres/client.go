package postgres

import (
	"database/sql"
	"fmt"

	"sbebe.ch/baby-name-guesser/pkg/utils"
)

type PostgresClient struct {
	db *sql.DB
}

func NewPostgresClient() (*PostgresClient, error) {
	connString := utils.GetDatabaseConnectionString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	return &PostgresClient{db: db}, nil
}

// Close closes the connection to the database.
func (c *PostgresClient) Close() error {
	return c.db.Close()
}

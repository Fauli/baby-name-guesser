package postgres

import (
	"fmt"

	_ "github.com/lib/pq"
	"sbebe.ch/baby-name-guesser/pkg/utils"
)

// GetAllBabyNames returns all baby names from the database ordered by alphabet.
func (c *PostgresClient) GetAllBabyNames() ([]string, error) {
	rows, err := c.db.Query("SELECT name FROM names ORDER BY name ASC")
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var babyNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to scan row: %v", err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		babyNames = append(babyNames, name)
	}

	return babyNames, nil
}

// AddBabyName adds a baby name to the database.
func (c *PostgresClient) AddBabyName(name string) error {
	utils.Logger.Sugar().Infof("Adding name to DB: %v\n", name)
	_, err := c.db.Exec("INSERT INTO names (name) VALUES ($1)", name)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// AddBabyNames adds a list of baby names to the database.
func (c *PostgresClient) AddBabyNames(names []string) error {
	utils.Logger.Sugar().Infof("Adding names to DB: %v\n", names)
	for _, name := range names {
		err := c.AddBabyName(name)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to add name: %v", err)
			return fmt.Errorf("failed to add name: %v", err)
		}
	}

	return nil
}

// DeleteBabyName deletes a baby name from the database.
func (c *PostgresClient) DeleteBabyName(name string) error {
	utils.Logger.Sugar().Infof("Deleting name from DB: %v\n", name)
	_, err := c.db.Exec("DELETE FROM names WHERE name = $1", name)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

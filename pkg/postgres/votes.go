package postgres

import (
	"fmt"

	"sbebe.ch/baby-name-guesser/pkg/utils"
)

// GetVoterVotesForNames returns a list of names and the number of votes for each name.
func (c *PostgresClient) GetVotesForNames() (map[string]int, error) {
	rows, err := c.db.Query("SELECT names.name, COUNT(votes.names_fk) FROM names LEFT JOIN votes ON names.name = votes.names_fk GROUP BY names.name")
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	votes := make(map[string]int)
	for rows.Next() {
		var name string
		var count int
		err := rows.Scan(&name, &count)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to scan row: %v", err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		votes[name] = count
	}

	return votes, nil

}

// get all votes returns a list of all emails and votes.
func (c *PostgresClient) GetAllVotesPerMail() (map[string][]string, error) {
	rows, err := c.db.Query("SELECT voter_fk, names_fk FROM votes")
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	votes := make(map[string][]string)
	for rows.Next() {
		var email string
		var name string
		err := rows.Scan(&email, &name)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to scan row: %v", err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		votes[email] = append(votes[email], name)
	}

	return votes, nil
}

// AddVotes adds a vote for a list of name.
func (c *PostgresClient) AddVote(email string, name string) error {

	fmt.Printf("Adding name to DB: %s, %s\n", email, name)
	_, err := c.db.Exec("INSERT INTO votes (names_fk, voter_fk) VALUES ($1, $2)", name, email)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// AddVotes adds a vote for a list of names.
func (c *PostgresClient) AddVotes(email string, names []string) error {
	fmt.Printf("Adding votes to DB for %s: %v\n", email, names)
	for _, name := range names {
		err := c.AddVote(email, name)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to add name: %v", err)
			return fmt.Errorf("failed to add name: %v", err)
		}
	}

	return nil
}

// DeleteVote deletes a vote from the list.
func (c *PostgresClient) DeleteVote(email string, name string) error {
	fmt.Printf("Deleting vote from DB for %s: %v\n", email, name)
	_, err := c.db.Exec("DELETE FROM votes WHERE names_fk = $1 AND voter_fk = $2", name, email)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// GetTopVotes returns the top 10 names with the most votes.
func (c *PostgresClient) GetTopVotes(limit int) (map[string]int, error) {
	rows, err := c.db.Query("SELECT names.name, COUNT(votes.names_fk) FROM names LEFT JOIN votes ON names.name = votes.names_fk GROUP BY names.name ORDER BY COUNT(votes.names_fk) DESC LIMIT $1", limit)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	votes := make(map[string]int)
	for rows.Next() {
		var name string
		var count int
		err := rows.Scan(&name, &count)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to scan row: %v", err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		votes[name] = count
	}

	return votes, nil
}

// GetVotesForVoter returns the votes for a voter.
func (c *PostgresClient) GetVotesForVoter(email string) ([]string, error) {
	rows, err := c.db.Query("SELECT names_fk FROM votes WHERE voter_fk = $1 ORDER BY names_fk ASC", email)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var votes []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to scan row: %v", err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		votes = append(votes, name)
	}

	return votes, nil
}

package postgres

import "fmt"

// GetVoterVotesForNames returns a list of names and the number of votes for each name.
func (c *PostgresClient) GetVotesForNames() (map[string]int, error) {
	rows, err := c.db.Query("SELECT names.name, COUNT(votes.names_fk) FROM names LEFT JOIN votes ON names.name = votes.names_fk GROUP BY names.name")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	votes := make(map[string]int)
	for rows.Next() {
		var name string
		var count int
		err := rows.Scan(&name, &count)
		if err != nil {
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
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	votes := make(map[string][]string)
	for rows.Next() {
		var email string
		var name string
		err := rows.Scan(&email, &name)
		if err != nil {
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
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// AddVotes adds a vote for a list of names.
func (c *PostgresClient) AddVotes(email string, names []string) error {
	fmt.Printf("Adding names to DB: %v\n", names)
	for _, name := range names {
		err := c.AddVote(email, name)
		if err != nil {
			return fmt.Errorf("failed to add name: %v", err)
		}
	}

	return nil
}

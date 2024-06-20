package postgres

import (
	"fmt"

	"sbebe.ch/baby-name-guesser/pkg/utils"
)

type VoterRow struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	HasVoted bool   `json:"has_voted"`
	HasPaid  bool   `json:"has_paid"`
}

// Add Voter inserts a single voter entry into the voters table
func (c *PostgresClient) AddVoter(name, lastName, email, passwort string) error {
	utils.Logger.Sugar().Infof("Adding voter to DB: %v\n", email)
	_, err := c.db.Exec("INSERT INTO voters (name, last_name, email, password) VALUES ($1, $2, $3, crypt($4, gen_salt('bf')))",
		name, lastName, email, passwort)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// GetAllVoters returns all voters from the voters table, joins the votes table to check if the voter has voted
func (c *PostgresClient) GetAllVoters() ([]VoterRow, error) {
	utils.Logger.Sugar().Info("Getting all voters from DB")
	// Select name, last_name, email from voters, joins votes to see if the voter has voted and paid
	rows, err := c.db.Query("SELECT v.name, v.last_name, v.email, COALESCE(vt.is_paid, false) AS is_paid, (vt.voter_fk IS NOT NULL) AS has_voted FROM public.voters v LEFT JOIN public.votes vt ON v.email = vt.voter_fk GROUP BY v.name, v.last_name, v.email, vt.is_paid, vt.voter_fk;")
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var voters []VoterRow
	for rows.Next() {
		var voter VoterRow
		err := rows.Scan(&voter.Name, &voter.LastName, &voter.Email, &voter.HasPaid, &voter.HasVoted)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to scan row: %v", err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		voters = append(voters, voter)
	}

	return voters, nil
}

// DeleteVoter deletes a single voter entry from the voters table
func (c *PostgresClient) DeleteVoter(email string) error {
	utils.Logger.Sugar().Infof("Deleting voter from DB: %v\n", email)
	_, err := c.db.Exec("DELETE FROM voters WHERE email = $1", email)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// LoginVoter compares the email address and passwort with the one in the database
func (c *PostgresClient) LoginVoter(email, passwort string) (bool, error) {
	utils.Logger.Sugar().Infof("Logging in voter: %v\n", email)
	//                     select email from voters where email = 'a' and password = crypt('b', password)
	rows, err := c.db.Query("SELECT email FROM voters WHERE email = $1 AND password = crypt($2, password)", email, passwort)
	if err != nil {
		utils.Logger.Sugar().Warnf("failed to execute query: %v\n", err)
		return false, fmt.Errorf("failed to execute query: %v", err)
	}

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			utils.Logger.Sugar().Warnf("Login failed: %v\n", email)
			return false, fmt.Errorf("failed to scan row: %v", err)
		}
	}

	if name == "" {
		utils.Logger.Sugar().Warnf("Login failed: %v\n", email)
		return false, fmt.Errorf("login failed")
	}

	utils.Logger.Sugar().Infof("Login successful: %v\n", email)
	return true, nil
}

// GetNameAndLastname returns the name and last name of a voter by email
func (c *PostgresClient) GetNameAndLastname(email string) (string, string, error) {
	fmt.Printf("Getting voter details by email: %v\n", email)
	rows, err := c.db.Query("SELECT name, last_name FROM voters WHERE email = $1", email)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return "", "", fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var name string
	var lastName string
	for rows.Next() {
		err := rows.Scan(&name, &lastName)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to scan row: %v", err)
			return "", "", fmt.Errorf("failed to scan row: %v", err)
		}
	}

	if name == "" || lastName == "" {
		return "", "", fmt.Errorf("no voter found with email: %v", email)
	}

	return name, lastName, nil
}

// HasUserVoted checks if a user has already voted
func (c *PostgresClient) HasUserVoted(email string) (bool, error) {
	fmt.Printf("Checking if user has voted: %v\n", email)
	rows, err := c.db.Query("SELECT voter_fk FROM votes WHERE voter_fk = $1", email)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return false, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var voter string
	for rows.Next() {
		err := rows.Scan(&voter)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to scan row: %v", err)
			return false, fmt.Errorf("failed to scan row: %v", err)
		}
	}

	if voter == "" {
		return false, nil
	}

	return true, nil
}

// PayForVotes sets the is_paid flag to true for a given email.
func (c *PostgresClient) PayForVotes(email string) error {
	fmt.Printf("Paying for votes: %v\n", email)
	_, err := c.db.Exec("UPDATE votes SET is_paid = true WHERE voter_fk = $1", email)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// HasUserPaid checks if a user has paid for their votes.
// Basically returns the row with the value of is_paid column.
func (c *PostgresClient) HasUserPaid(email string) (bool, error) {
	fmt.Printf("Checking if user has paid: %v\n", email)
	rows, err := c.db.Query("SELECT is_paid FROM votes WHERE voter_fk = $1 AND is_paid = true", email)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to execute query: %v", err)
		return false, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	isPaid := false
	for rows.Next() {
		err := rows.Scan(&isPaid)
		if err != nil {
			utils.Logger.Sugar().Errorf("failed to scan row: %v", err)
			return false, fmt.Errorf("failed to scan row: %v", err)
		}
	}

	return isPaid, nil
}

// CREATE EXTENSION pgcrypto;

// CREATE TABLE IF NOT EXISTS public.names
// (
//     name character varying COLLATE pg_catalog."default" NOT NULL,
//     CONSTRAINT names_pkey PRIMARY KEY (name)
// )

// CREATE TABLE IF NOT EXISTS public.voters
// (
//     name character varying COLLATE pg_catalog."default" NOT NULL,
//     last_name character varying COLLATE pg_catalog."default" NOT NULL,
//     email character varying COLLATE pg_catalog."default" NOT NULL,
//     password character varying COLLATE pg_catalog."default" NOT NULL,
//     CONSTRAINT voters_pkey PRIMARY KEY (email)
// )

// INSERT INTO public.voters(
// 	name, last_name, email, password)
// 	VALUES ('Franz', 'Faul', 'f@sbebe.ch',
// 			crypt('franzpassword', gen_salt('bf'))
// 			);

// SELECT id
//   FROM users
//  WHERE email = 'johndoe@mail.com'
//    AND password = crypt('johnspassword', password);

// CREATE TABLE IF NOT EXISTS public.voters
// (
//     name character varying COLLATE pg_catalog."default" NOT NULL,
//     last_name character varying COLLATE pg_catalog."default" NOT NULL,
//     email character varying COLLATE pg_catalog."default" NOT NULL,
//     password character varying COLLATE pg_catalog."default" NOT NULL,
//     CONSTRAINT voters_pkey PRIMARY KEY (email)
// )

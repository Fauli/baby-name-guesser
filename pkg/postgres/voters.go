package postgres

import (
	"fmt"
)

// Add Voter inserts a single voter entry into the voters table
func (c *PostgresClient) AddVoter(name, lastName, email, passwort string) error {
	fmt.Printf("Adding voter to DB: %v\n", email)
	_, err := c.db.Exec("INSERT INTO voters (name, last_name, email, password) VALUES ($1, $2, $3, crypt($4, gen_salt('bf')))",
		name, lastName, email, passwort)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// DeleteVoter deletes a single voter entry from the voters table
func (c *PostgresClient) DeleteVoter(email string) error {
	fmt.Printf("Deleting voter from DB: %v\n", email)
	_, err := c.db.Exec("DELETE FROM voters WHERE email = $1", email)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// LoginVoter compares the email address and passwort with the one in the database
func (c *PostgresClient) LoginVoter(email, passwort string) (bool, error) {
	fmt.Printf("Logging in voter: %v\n", email)
	//                     select email from voters where email = 'a' and password = crypt('b', password)
	rows, err := c.db.Query("SELECT email FROM voters WHERE email = $1 AND password = crypt($2, password)", email, passwort)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %v", err)
	}

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			return false, fmt.Errorf("failed to scan row: %v", err)
		}
	}

	if name == "" {
		return false, fmt.Errorf("login failed")
	}

	fmt.Println("Login successful")
	fmt.Println(name)
	return true, nil
}

// GetNameAndLastname returns the name and last name of a voter by email
func (c *PostgresClient) GetNameAndLastname(email string) (string, string, error) {
	fmt.Printf("Getting voter details by email: %v\n", email)
	rows, err := c.db.Query("SELECT name, last_name FROM voters WHERE email = $1", email)
	if err != nil {
		return "", "", fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var name string
	var lastName string
	for rows.Next() {
		err := rows.Scan(&name, &lastName)
		if err != nil {
			return "", "", fmt.Errorf("failed to scan row: %v", err)
		}
	}

	if name == "" || lastName == "" {
		return "", "", fmt.Errorf("no voter found with email: %v", email)
	}

	return name, lastName, nil
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

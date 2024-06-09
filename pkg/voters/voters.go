package voters

import (
	"fmt"

	"sbebe.ch/baby-name-guesser/pkg/postgres"
	"sbebe.ch/baby-name-guesser/pkg/utils"
)

type Voter struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	HasVoted bool   `json:"has_voted"`
}

type VoterFull struct {
	Voter
	Password      string `json:"password"`
	EventPassword string `json:"event_password"`
}

type VoterLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddVoter(voter VoterFull) (Voter, error) {

	// Validate the input
	if voter.Name == "" || voter.LastName == "" || voter.Email == "" || voter.Password == "" {
		return Voter{}, fmt.Errorf("missing fields")
	}

	if len(voter.Password) < 8 {
		return Voter{}, fmt.Errorf("password too short")
	}

	if !utils.IsMailValid(voter.Email) {
		return Voter{}, fmt.Errorf("invalid email")
	}

	if voter.EventPassword != utils.GetEventPassword() {
		return Voter{}, fmt.Errorf("invalid event password")
	}

	// Add the voter
	fmt.Printf("Adding new voter: %s %s [%s]\n", voter.Name, voter.LastName, voter.Email)
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return Voter{}, err
	}
	defer c.Close()
	err = c.AddVoter(voter.Name, voter.LastName, voter.Email, voter.Password)
	if err != nil {
		return Voter{}, err

	}

	return Voter{Email: voter.Email, Name: voter.Name, LastName: voter.LastName}, nil
}

func DeleteVoterByEmail(email string) error {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return err
	}
	defer c.Close()
	err = c.DeleteVoter(email)
	return err
}

func GetVoterByEmail(email string) (Voter, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return Voter{}, err
	}
	defer c.Close()
	name, lastName, err := c.GetNameAndLastname(email)
	if err != nil {
		return Voter{}, err
	}
	return Voter{Email: email, Name: name, LastName: lastName}, nil
}

func LoginVoter(email, password string) (bool, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return false, err
	}
	defer c.Close()
	result, err := c.LoginVoter(email, password)
	return result, err
}

func GetNameAndLastname(email string) (string, string, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return "", "", err
	}
	defer c.Close()
	name, lastName, err := c.GetNameAndLastname(email)
	return name, lastName, err
}

func HasUserVoted(email string) (bool, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return false, err
	}
	defer c.Close()
	return c.HasUserVoted(email)
}

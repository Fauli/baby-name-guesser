package voters

import (
	"fmt"

	"sbebe.ch/baby-name-guesser/pkg/postgres"
)

type Voter struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

type VoterFull struct {
	Voter
	Password string `json:"password"`
}

type VoterLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddVoter(voter VoterFull) (Voter, error) {
	fmt.Println("adding names")
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return Voter{}, err
	}
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
	err = c.DeleteVoter(email)
	return err
}

func GetVoterByEmail(email string) (Voter, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return Voter{}, err
	}
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
	result, err := c.LoginVoter(email, password)
	return result, err
}

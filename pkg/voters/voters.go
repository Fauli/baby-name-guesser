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
	HasPaid  bool   `json:"has_paid"`
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
		utils.Logger.Sugar().Errorf("failed to create postgres client: %v", err)
		return Voter{}, err
	}
	defer c.Close()
	err = c.AddVoter(voter.Name, voter.LastName, voter.Email, voter.Password)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to add voter: %v", err)
		return Voter{}, err

	}

	return Voter{Email: voter.Email, Name: voter.Name, LastName: voter.LastName}, nil
}

func DeleteVoterByEmail(email string) error {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to create postgres client: %v", err)
		return err
	}
	defer c.Close()
	err = c.DeleteVoter(email)
	return err
}

func GetAllVoters() ([]Voter, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to create postgres client: %v", err)
		return nil, err
	}
	defer c.Close()
	voterRows, err := c.GetAllVoters()
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to get voters: %v", err)
		return nil, err
	}

	// TODO: [franz] this is a bit ugly, fix it please
	voters := make([]Voter, len(voterRows))
	for i, voterRow := range voterRows {
		voters[i] = Voter{
			Name:     voterRow.Name,
			LastName: voterRow.LastName,
			Email:    voterRow.Email,
		}
	}
	return voters, nil
}

func GetVoterByEmail(email string) (Voter, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to create postgres client: %v", err)
		return Voter{}, err
	}
	defer c.Close()
	name, lastName, err := c.GetNameAndLastname(email)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to get voter: %v", err)
		return Voter{}, err
	}
	return Voter{Email: email, Name: name, LastName: lastName}, nil
}

func LoginVoter(email, password string) (bool, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to create postgres client: %v", err)
		return false, err
	}
	defer c.Close()
	result, err := c.LoginVoter(email, password)
	return result, err
}

func GetNameAndLastname(email string) (string, string, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to create postgres client: %v", err)
		return "", "", err
	}
	defer c.Close()
	name, lastName, err := c.GetNameAndLastname(email)
	return name, lastName, err
}

func HasUserVoted(email string) (bool, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to create postgres client: %v", err)
		return false, err
	}
	defer c.Close()
	return c.HasUserVoted(email)
}

// PayForVotes pays for the votes.
func PayForVotes(email string) error {
	// Validate the input
	if email == "" {
		return fmt.Errorf("no email provided")
	}

	// Pay for the votes
	c, err := postgres.NewPostgresClient()
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to create postgres client: %v", err)
		return err
	}
	defer c.Close()
	err = c.PayForVotes(email)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to pay for votes: %v", err)
		return err
	}
	return nil
}

// HasUserPaid checks if a user has already paid
func HasUserPaid(email string) (bool, error) {
	// Validate the input
	if email == "" {
		return false, fmt.Errorf("no email provided")
	}

	// Check if the user has paid
	c, err := postgres.NewPostgresClient()
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to create postgres client: %v", err)
		return false, err
	}
	defer c.Close()
	return c.HasUserPaid(email)
}

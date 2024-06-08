package votes

import (
	"fmt"

	"sbebe.ch/baby-name-guesser/pkg/postgres"
)

const (
	topVotes = 3
)

type Vote struct {
	Email string   `json:"email"`
	Names []string `json:"names"`
}

func AddVotes(voter string, votes []string) (Vote, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return Vote{}, err
	}
	defer c.Close()
	err = c.AddVotes(voter, votes)
	if err != nil {
		return Vote{}, err

	}

	return Vote{Email: voter, Names: votes}, nil
}

func GetVotes() (map[string]int, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.GetVotesForNames()
}

func GetVotesForVoters() ([]Vote, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	voteMap, err := c.GetAllVotesPerMail()
	if err != nil {
		return nil, err
	}

	result := make([]Vote, 0)
	for mail, votes := range voteMap {
		fmt.Printf("Mail: %s, Votes: %v\n", mail, votes)
		result = append(result, Vote{Email: mail, Names: votes})
	}

	return result, nil

}

func GetTopVotes() (map[string]int, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.GetTopVotes(topVotes)
}

package votes

import (
	"fmt"

	"sbebe.ch/baby-name-guesser/pkg/postgres"
)

type Vote struct {
	Email string   `json:"email"`
	Names []string `json:"names"`
}

func AddVotes(vote Vote) (Vote, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return Vote{}, err
	}
	err = c.AddVotes(vote.Email, vote.Names)
	if err != nil {
		return Vote{}, err

	}

	return Vote{Email: vote.Email, Names: vote.Names}, nil
}

func GetVotes() (map[string]int, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return nil, err
	}
	return c.GetVotesForNames()
}

func GetVotesForVoters() ([]Vote, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return nil, err
	}
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

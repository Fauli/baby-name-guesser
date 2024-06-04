package votes

import "sbebe.ch/baby-name-guesser/pkg/postgres"

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

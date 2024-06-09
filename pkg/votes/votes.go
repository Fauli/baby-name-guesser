package votes

import (
	"fmt"
	"sort"

	"sbebe.ch/baby-name-guesser/pkg/postgres"
	"sbebe.ch/baby-name-guesser/pkg/utils"
	"sbebe.ch/baby-name-guesser/pkg/voters"
)

const (
	topVotes = 3
)

type Vote struct {
	Email string   `json:"email"`
	Names []string `json:"names"`
}

func AddVotes(email string, votes []string) (Vote, error) {

	// Validate the input
	if len(votes) == 0 {
		return Vote{}, fmt.Errorf("no names to add")
	}

	hasVoted, err := voters.HasUserVoted(email)
	if err != nil {
		return Vote{}, fmt.Errorf("failed to check if user has voted: %v", err)
	}

	if hasVoted {
		return Vote{}, fmt.Errorf("user has already voted")
	}

	// Add the votes
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return Vote{}, err
	}
	defer c.Close()
	err = c.AddVotes(email, votes)
	if err != nil {
		return Vote{}, err

	}

	return Vote{Email: email, Names: votes}, nil
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

func GetTopVotes() ([]utils.KV, error) {
	// TODO: [franz] ensure the list is properly sorted when returned
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	votes, err := c.GetTopVotes(topVotes)
	if err != nil {
		return nil, err
	}

	vec := utils.MapToSlice(votes)

	sort.Slice(vec, func(i, j int) bool {
		// 1. value is different - sort by value (in reverse order)
		if vec[i].Value != vec[j].Value {
			return vec[i].Value > vec[j].Value
		}
		// 2. only when value is the same - sort by key
		return vec[i].Key < vec[j].Key
	})

	fmt.Printf("%v", vec)

	return vec, nil
}

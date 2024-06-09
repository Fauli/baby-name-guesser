package utils

import (
	"net/mail"
)

func IsMailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func MapToSlice(in map[string]int) []KV {
	vec := make([]KV, len(in))
	i := 0
	for k, v := range in {
		vec[i].Key = k
		vec[i].Value = v
		i++
	}
	return vec
}

type KV struct {
	Key   string
	Value int
}

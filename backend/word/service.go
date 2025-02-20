package word

import (
	"fmt"
	"math/rand/v2"
)

type WordService struct {
	words     []string
	today     string
	yesterday string
}

func NewWordService(words []string) *WordService {
	s := &WordService{
		words:     words,
		today:     "",
		yesterday: "",
	}

	s.today = s.words[rand.IntN(len(s.words))]
	s.yesterday = s.today
	s.today = s.words[rand.IntN(len(s.words))]

	return s
}

// TODO: change this every night at midnight
func (s *WordService) ChangeWords() (string, string) {
	s.yesterday = s.today
	s.today = s.words[rand.IntN(len(s.words))]
	return s.yesterday, s.today
}

type LetterCheck struct {
	Letter string `json:"letter"`
	Status string `json:"status"`
}

// TODO: write unit tests to find errors, because this isn't quite working yet
func (s *WordService) CheckGuess(guess string) ([]LetterCheck, error) {
	// if guess not in allowed

	// if guess not the proper length
	if len(guess) != len(s.today) {
		return nil, fmt.Errorf("guess must match the length of the answer")
	}

	results := make([]LetterCheck, len(guess))

	matched := make([]bool, len(s.today))

	// check exact matches
	for i, char := range guess {
		results[i].Letter = string(char)

		if char == rune(s.today[i]) {
			results[i].Status = "correct"
			matched[i] = true
		}

		for i, guessChar := range guess {
			if results[i].Status == "correct" {
				continue
			}

			found := false
			for j, answerChar := range s.today {
				if !matched[j] && guessChar == answerChar {
					results[i].Status = "present"
					matched[j] = true
					found = true
					break
				}
			}

			if !found {
				results[i].Status = "absent"
			}
		}
	}

	return results, nil
}

package word

import (
	"math/rand/v2"
)

type WordService struct {
	words     []string
	today     int
	yesterday int
}

func NewWordService(words []string) *WordService {
	// TODO: seed random number
	// TODO: read yesterday from persistance
	return &WordService{
		words:     words,
		today:     rand.IntN(len(words)),
		yesterday: 0,
	}
}

func (s *WordService) GetWords() (string, string) {
	return s.words[s.yesterday], s.words[s.today]
}

func (s *WordService) ChangeWords() (string, string) {
	s.yesterday = s.today
	s.today = rand.IntN(len(s.words))

	return s.words[s.yesterday], s.words[s.today]
}

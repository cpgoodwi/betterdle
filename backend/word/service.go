package word

import "math/rand/v2"

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
func (s WordService) ChangeWords() (string, string) {
	s.yesterday = s.today
	s.today = s.words[rand.IntN(len(s.words))]
	return s.yesterday, s.today
}

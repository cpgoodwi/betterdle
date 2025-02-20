package word

import (
	"encoding/json"
	"net/http"
	"strings"
)

type WordHandler struct {
	s WordService
}

func NewWordHandler(s WordService) *WordHandler {
	return &WordHandler{
		s: s,
	}
}

type WordResponse struct {
	Yesterday string `json:"yesterday"`
	Today     string `json:"today"`
	Error     string `json:"error"`
}

func (h *WordHandler) GetWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(WordResponse{
		Yesterday: h.s.yesterday,
		Today:     h.s.today,
	})
}

func (h *WordHandler) ChangeWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	yesterday, today := h.s.ChangeWords()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(WordResponse{
		Yesterday: yesterday,
		Today:     today,
	})
}

type GuessRequest struct {
	Guess string `json:"guess"`
}

type GuessResponse struct {
	Results []LetterCheck `json:"results"`
	Error   string        `json:"error"`
}

func (h *WordHandler) CheckGuess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// decode request body
	var req GuessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GuessResponse{
			Error: "invalid request body",
		})
		return
	}

	// validate guess
	req.Guess = strings.TrimSpace(strings.ToLower(req.Guess))
	if req.Guess == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GuessResponse{
			Error: "guess cannot be empty",
		})
		return
	}

	// check the guess
	results, err := h.s.CheckGuess(req.Guess)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GuessResponse{
			Error: err.Error(),
		})
		return
	}

	// return successful response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GuessResponse{
		Results: results,
	})
}

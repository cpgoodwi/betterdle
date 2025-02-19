package word

import (
	"encoding/json"
	"net/http"
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

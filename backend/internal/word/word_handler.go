package word

import (
	"fmt"
	"net/http"
)

type WordHandler struct {
	service WordService
}

func NewWordHandler(service WordService) *WordHandler {
	return &WordHandler{
		service: service,
	}
}

func (h *WordHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// indexStr := r.URL.Query().Get("index")
	// if indexStr == "" {
	// 	http.Error(w, "please provide a word index number using ?index=N", http.StatusBadRequest)
	// 	return
	// }

	// index, err := strconv.Atoi(indexStr)
	// if err != nil {
	// 	http.Error(w, "invalid index", http.StatusBadRequest)
	// 	return
	// }

	// word, err := h.service.GetWordFromLine(index)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// }

	yesterday, today := h.service.ChangeWords()

	fmt.Fprint(w, yesterday, today)
}

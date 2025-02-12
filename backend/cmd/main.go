package main

import (
	"betterdle-backend/internal/word"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	wordService := word.NewWordService("data/words.txt")
	wordHandler := word.NewWordHandler(*wordService)

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/word", wordHandler.Handle)

	fmt.Println("server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

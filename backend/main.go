package main

import (
	"betterdle-server/data"
	"betterdle-server/word"
	"fmt"
	"log"
	"net/http"
)

func main() {
	wordService := word.NewWordService(data.Words)
	wordHandler := word.NewWordHandler(*wordService)

	router := http.NewServeMux()

	router.HandleFunc("GET /word", wordHandler.GetWords)
	router.HandleFunc("PUT /word", wordHandler.ChangeWords)
	router.HandleFunc("POST /word", wordHandler.CheckGuess)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is starting on port 8080")

	log.Fatal(server.ListenAndServe())
}

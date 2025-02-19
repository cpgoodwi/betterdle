package main

import (
	"betterdle-server/data"
	"betterdle-server/word"
	"fmt"
	"net/http"
)

func main() {
	wordService := word.NewWordService(data.Words)
	wordHandler := word.NewWordHandler(*wordService)

	router := http.NewServeMux()

	router.HandleFunc("GET /word", wordHandler.GetWords)
	router.HandleFunc("PUT /word", wordHandler.ChangeWords)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is starting on port 8080")
	server.ListenAndServe()
}

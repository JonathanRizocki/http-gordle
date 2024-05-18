package main

import (
	"learngo/httpgordle/internal/handlers"
	"learngo/httpgordle/internal/repository"
	"net/http"
)

func main() {
	// Start the server
	db := repository.New()

	err := http.ListenAndServe(":8080", handlers.NewRouter(db))
	if err != nil {
		panic(err)
	}
}

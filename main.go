package main

import (
	"learngo/httpgordle/internal/handlers"
	"net/http"
)

func main() {
	// Start the server
	err := http.ListenAndServe(":8080", handlers.NewRouter())
	if err != nil {
		panic(err)
	}
}

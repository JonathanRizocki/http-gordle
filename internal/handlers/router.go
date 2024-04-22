package handlers

import (
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/handlers/newgame"

	"github.com/go-chi/chi/v5"
)

// NewRouter turns a router that listens for requests
// to the following endpoints:
// - Create a new game;
//
// The provided router is ready to serve
func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Post(api.NewGameRoute, newgame.Handle)

	return r
}

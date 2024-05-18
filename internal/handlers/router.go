package handlers

import (
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/handlers/getstatus"
	"learngo/httpgordle/internal/handlers/guess"
	"learngo/httpgordle/internal/handlers/newgame"
	"learngo/httpgordle/internal/repository"

	"github.com/go-chi/chi/v5"
)

// NewRouter turns a router that listens for requests
// to the following endpoints:
// - Create a new game;
//
// The provided router is ready to serve
func NewRouter(db *repository.GameRepository) chi.Router {
	r := chi.NewRouter()

	r.Post(api.NewGameRoute, newgame.Handler(db))
	r.Get(api.GetStatusRoute, getstatus.Handler(db))
	r.Put(api.GuessRoute, guess.Handler(db))

	return r
}

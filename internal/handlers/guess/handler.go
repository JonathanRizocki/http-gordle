package guess

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type gameGuesser interface {
	Find(session.GameID) (session.Game, error)
	Update(session.GameID, session.Game) (session.Game, error)
}

// Handler returns the handler for the guess endpoint.
func Handler(db gameGuesser) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, api.GameID)

		if id == "" {
			http.Error(w, "missing the id of the game", http.StatusNotFound)
			return
		}

		r := api.Guess{}
		err := json.NewDecoder(req.Body).Decode(&r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		game, err := guess(id, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}
	}
}

func guess(id string, r api.Guess) (session.Game, error) {
	return session.Game{
		ID: session.GameID(id),
	}, nil
}

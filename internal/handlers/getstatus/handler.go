package getstatus

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handler returns the handler for the game status endpoint.
// The repo parameter will be more clearly defined in the next section.
func Handler(repo interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, api.GameID)

		if id == "" {
			http.Error(w, "missing the id of the game", http.StatusBadRequest)
			return
		}

		game := getGame(id)

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			// The header has already been set.
			// Nothing else we can do here.
			log.Printf("failed to write response: %s", err)
		}

	}
}

func getGame(id string) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}

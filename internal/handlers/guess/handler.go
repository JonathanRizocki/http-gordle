package guess

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Handle(w http.ResponseWriter, req *http.Request) {
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

	apiGame := api.GameResponse{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		// The header has already been set. Nothing much we can do here.
		log.Printf("failed to write response: %s", err)
	}
}

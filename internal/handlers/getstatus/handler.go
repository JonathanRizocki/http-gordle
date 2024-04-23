package getstatus

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
		http.Error(w, "missing the id of the game", http.StatusBadRequest)
		return
	}
	log.Printf("retrive status of game with id: %v", id)

	apiGame := api.GameResponse{
		ID: id,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		// The header has already been set.
		// Nothing else we can do here.
		log.Printf("failed to write response: %s", err)
	}

}

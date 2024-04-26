package repository

import (
	"learngo/httpgordle/internal/session"
)

// GameRepository holds all the current games.
type GameRepository struct {
	storage map[session.GameID]session.Game
}

// New creates an empty game repository.
func New() *GameRepository {
	return &GameRepository{
		storage: make(map[session.GameID]session.Game),
	}
}

// Add inserts for the first time a game in memory.
func (gr *GameRepository) Add(game session.Game) error {
	_, ok := gr.storage[game.ID]
	if ok {
		return ErrConflictingId
	}

	gr.storage[game.ID] = game
	return nil
}

// Find retrieves a game if it exists in memory.
func (gr *GameRepository) Find(gameId session.GameID) (session.Game, error) {
	game, ok := gr.storage[gameId]

	if !ok {
		game = session.Game{}
		return game, ErrNotFound
	}

	return game, nil
}

// Update retrieves a game if it exists and modifies its state
func (gr *GameRepository) Update(gameId session.GameID, update session.Game) (session.Game, error) {
	game, ok := gr.storage[gameId]

	emptyGame := session.Game{}

	if !ok {
		return emptyGame, ErrNotFound
	}

	if game.Status != session.StatusPlaying {
		return emptyGame, ErrGameNotActive
	}

	gr.storage[gameId] = update
	game = gr.storage[gameId]

	return game, nil
}

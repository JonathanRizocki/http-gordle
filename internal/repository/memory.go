package repository

import (
	"fmt"
	"learngo/httpgordle/internal/session"
	"log"
	"sync"
)

// GameRepository holds all the current games.
type GameRepository struct {
	mutex   sync.Mutex
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
	log.Print("Adding a game...")

	// Lock the reading and writing of the game
	gr.mutex.Lock()
	defer gr.mutex.Unlock()

	_, ok := gr.storage[game.ID]
	if ok {
		return fmt.Errorf("%w (%s)", ErrConflictingID, game.ID)
	}

	gr.storage[game.ID] = game
	return nil
}

// Find retrieves a game if it exists in memory.
func (gr *GameRepository) Find(gameId session.GameID) (session.Game, error) {
	log.Print("Finding a game...")
	game, ok := gr.storage[gameId]

	// Lock the reading and writing of the game
	gr.mutex.Lock()
	defer gr.mutex.Unlock()

	if !ok {
		game = session.Game{}
		return game, fmt.Errorf("%w (%s)", ErrNotFound, game.ID)
	}

	return game, nil
}

// Update retrieves a game if it exists and modifies its state
func (gr *GameRepository) Update(gameId session.GameID, update session.Game) (session.Game, error) {
	log.Print("Updating a game...")

	// Lock the reading and writing of the game
	gr.mutex.Lock()
	defer gr.mutex.Unlock()

	game, ok := gr.storage[gameId]

	emptyGame := session.Game{}

	if !ok {
		return emptyGame, fmt.Errorf("%w (%s)", ErrNotFound, game.ID)
	}

	if game.Status != session.StatusPlaying {
		return emptyGame, fmt.Errorf("%w (%s)", ErrGameNotActive, game.ID)
	}

	gr.storage[gameId] = update
	game = gr.storage[gameId]

	return game, nil
}

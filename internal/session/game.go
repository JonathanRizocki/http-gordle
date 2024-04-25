package session

import (
	"errors"
)

// Game contains the information about a game.
type Game struct {
	ID           GameID
	AttemptsLeft byte
	Guesses      []Guess
	Status       Status
}

// A GameID represents the ID of a game.
type GameID string

// Status is the current status of the game and
// tells what operations can be made on it.
type Status string

// Guess is a pair of a word (submitted by the player) and its feedback
type Guess struct {
	Word     string
	Feedback string
}

const (
	StatusPlaying = "Playing"
	StatusWon     = "Won"
	StatusLost    = "Lost"
)

// ErrGameOver is returned when a play is made by the game is over.
var ErrGameOver = errors.New("game over")

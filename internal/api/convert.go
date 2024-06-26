package api

import (
	"learngo/httpgordle/internal/session"
)

// ToGameResponse converts a session.Game into a GameResponse.
func ToGameResponse(g session.Game) GameResponse {
	apiGame := GameResponse{
		ID:           string(g.ID),
		AttemptsLeft: g.AttemptsLeft,
		Guesses:      make([]Guess, len(g.Guesses)),
		Status:       string(g.Status),
		// TODO wordlength
	}

	for i := 0; i < len(g.Guesses); i++ {
		apiGame.Guesses[i].Word = g.Guesses[i].Word
		apiGame.Guesses[i].Feedback = g.Guesses[i].Feedback
	}

	if g.AttemptsLeft == 0 {
		apiGame.Solution = "" // TODO solution
	}

	return apiGame
}

package api

import (
	"learngo/httpgordle/internal/session"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	id := "12345678"
	tt := map[string]struct {
		game session.Game
		want GameResponse
	}{
		"nominal": {
			game: session.Game{
				ID:           session.GameID(id),
				AttemptsLeft: 4,
				Guesses: []session.Guess{
					{Word: "HELLO", Feedback: "⬜️🟡⬜️⬜️⬜️"},
				},
				Status: session.StatusPlaying,
			},
			want: GameResponse{
				ID:           id,
				AttemptsLeft: 4,
				Guesses: []Guess{
					{Word: "HELLO", Feedback: "⬜️🟡⬜️⬜️⬜️"},
				},
				Status: session.StatusPlaying,
			},
		},
		"multiple guesses": {
			game: session.Game{
				ID:           session.GameID(id),
				AttemptsLeft: 3,
				Guesses: []session.Guess{
					{Word: "HELLO", Feedback: "⬜️🟡⬜️⬜️⬜️"},
					{Word: "FENDS", Feedback: "⬜️🟡⬜️⬜️⬜️"},
				},
				Status: session.StatusPlaying,
			},
			want: GameResponse{
				ID:           id,
				AttemptsLeft: 3,
				Guesses: []Guess{
					{Word: "HELLO", Feedback: "⬜️🟡⬜️⬜️⬜️"},
					{Word: "FENDS", Feedback: "⬜️🟡⬜️⬜️⬜️"},
				},
				Status: session.StatusPlaying,
			},
		},
		"no guesses": {
			game: session.Game{
				ID:           session.GameID(id),
				AttemptsLeft: 5,
				Guesses:      []session.Guess{},
				Status:       session.StatusPlaying,
			},
			want: GameResponse{
				ID:           id,
				AttemptsLeft: 5,
				Guesses:      []Guess{},
				Status:       session.StatusPlaying,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := ToGameResponse(tc.game)
			assert.Equal(t, tc.want, got)
		})
	}
}

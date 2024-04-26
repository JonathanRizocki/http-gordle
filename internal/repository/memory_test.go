package repository

import (
	"learngo/httpgordle/internal/session"
	"testing"
)

const (
	gameID1 = "123456"
	gameID2 = "912372"
)

func TestNew(t *testing.T) {
	gameRepo := New()

	if gameRepo == nil {
		t.Error("game repo is nil")
	}
}

func TestAdd(t *testing.T) {
	gameRepo := New()
	newGame := session.Game{
		ID: gameID1,
	}

	err := gameRepo.Add(newGame)
	if err != nil {
		t.Error("unable to add new game")
	}

	newGame2 := session.Game{
		ID: gameID1,
	}

	err = gameRepo.Add(newGame2)
	if err.Error() != ErrConflictingId.Error() {
		t.Errorf("expected error %s, but got %s", ErrConflictingId, err)
	}

}

func TestFind(t *testing.T) {
	gameRepo := New()

	newGame := session.Game{
		ID: gameID1,
	}

	gameRepo.Add(newGame)

	game, err := gameRepo.Find(gameID1)
	if err != nil {
		t.Errorf("unable to find game")
	}
	if game.ID != gameID1 {
		t.Errorf("expected id %s, but found %s", gameID1, game.ID)
	}

	game, err = gameRepo.Find(gameID2)
	if err != ErrNotFound {
		t.Errorf("expected %v, but got %v", ErrNotFound, err.Error())
	}
}

func TestUpdate(t *testing.T) {
	gameRepo := New()

	newGame := session.Game{
		ID:           gameID1,
		AttemptsLeft: 1,
		Status:       session.StatusPlaying,
	}

	gameRepo.Add(newGame)

	gameUpdate := session.Game{
		ID:           gameID2,
		AttemptsLeft: 2,
	}

	result, err := gameRepo.Update(gameID2, gameUpdate)

	if err != ErrNotFound {
		t.Errorf("expected error %v, but got %v", ErrNotFound, err.Error())
	}
	if result.ID != "" {
		t.Errorf("expected empty struct")
	}

	gameUpdate = session.Game{
		ID:           gameID1,
		AttemptsLeft: 0,
	}

	result, err = gameRepo.Update(gameID1, gameUpdate)

	if err != nil {
		t.Errorf("expected no error, but got %v", err.Error())
	}

	if result.AttemptsLeft != 0 {
		t.Errorf("invalid update")
	}

}

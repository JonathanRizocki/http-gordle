package getstatus

import (
	"context"
	"learngo/httpgordle/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/games/", nil)
	require.NoError(t, err)

	// add path params
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(api.GameID, "123456")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	recorder := httptest.NewRecorder()

	Handle(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	wantedJson := `{"id":"123456","attempts_left":0,"guesses":null,"word_length":0,"status":""}`
	assert.JSONEq(t, wantedJson, recorder.Body.String())

}

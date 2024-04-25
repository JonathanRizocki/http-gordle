package guess

import (
	"context"
	"fmt"
	"learngo/httpgordle/internal/api"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {

	body := strings.NewReader(`{"word":"hello"}`)
	req, err := http.NewRequest(http.MethodPut, api.GuessRoute, body)
	require.NoError(t, err)

	// add path params
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(api.GameID, "123456")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	recorder := httptest.NewRecorder()

	Handle(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	wantedJson := `{"id":"123456","attempts_left":0,"guesses":[],"word_length":0,"status":""}`
	fmt.Printf("\nrecorder body %v\n", recorder.Body.String())
	assert.JSONEq(t, wantedJson, recorder.Body.String())
}

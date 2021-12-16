package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	server := &PlayerServer{}
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		responce := httptest.NewRecorder()

		server.ServerHTTP(responce, request)

		assertResponcseBody(t, responce.Body.String(), "20")

		//got := responce.Body.String()
		//want := "20"
		//
		//if got != want {
		//	t.Errorf("got %q, want %q", got, want)
		//}
	})

	t.Run("returns Floyd score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		responce := httptest.NewRecorder()

		server.ServerHTTP(responce, request)

		assertResponcseBody(t, responce.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponcseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

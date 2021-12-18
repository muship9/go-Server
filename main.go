package main

import (
	"log"
	"net/http"
)

type inMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *inMemoryPlayerStore {
	return &inMemoryPlayerStore{map[string]int{}}
}

func (i *inMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *inMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

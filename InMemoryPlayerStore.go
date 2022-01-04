package main

func NewInMemoryPlayerStore() *inMemoryPlayerStore {
	return &inMemoryPlayerStore{map[string]int{}}
}

type inMemoryPlayerStore struct {
	store map[string]int
}

func (i *inMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

func (i *inMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *inMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

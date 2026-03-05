package lobby

import (
	"sync"
)

type Lobby struct {
	ID        string     `json:"id"`
	HostID    string     `json:"host"`
	MaxPlayer int        `json:"max_players"`
	Players   []string   `json:"players"`
	Mutex     sync.Mutex `json:"-"`
}

func NewLobby(id string, host string, max int) *Lobby {

	return &Lobby{
		ID:        id,
		HostID:    host,
		MaxPlayer: max,
		Players:   []string{host},
	}
}

func (l *Lobby) AddPlayer(playerID string) bool {

	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	// check if player already exists
	for _, p := range l.Players {
		if p == playerID {
			return false
		}
	}

	if len(l.Players) >= l.MaxPlayer {
		return false
	}

	l.Players = append(l.Players, playerID)
	return true
}

func (l *Lobby) RemovePlayer(playerID string) {

	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	for i, p := range l.Players {

		if p == playerID {
			l.Players = append(l.Players[:i], l.Players[i+1:]...)
			break
		}
	}
}

package lobby

import (
	"sync"

	"github.com/google/uuid"
)

type Manager struct {
	Lobbies map[string]*Lobby
	Mutex   sync.Mutex
}

func NewManager() *Manager {

	return &Manager{
		Lobbies: make(map[string]*Lobby),
	}
}

func (m *Manager) CreateLobby(host string) *Lobby {

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	id := uuid.New().String()

	l := NewLobby(id, host, 4)

	m.Lobbies[id] = l

	return l
}

func (m *Manager) GetLobby(id string) (*Lobby, bool) {

	l, ok := m.Lobbies[id]
	return l, ok
}

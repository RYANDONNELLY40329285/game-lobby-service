package party

import "sync"

type Manager struct {
	Parties map[string]*Party
	Mutex   sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		Parties: make(map[string]*Party),
	}
}

func (m *Manager) CreateParty(leader string) *Party {

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	p := NewParty(leader)

	m.Parties[p.ID] = p

	return p
}

func (m *Manager) GetParty(id string) (*Party, bool) {

	p, ok := m.Parties[id]

	return p, ok
}

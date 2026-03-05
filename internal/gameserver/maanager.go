package gameserver

type GameServer struct {
	IP   string
	Port int
}

type Manager struct {
	servers []GameServer
	index   int
}

func NewManager() *Manager {
	return &Manager{
		servers: []GameServer{
			{"192.168.1.10", 7777},
			{"192.168.1.11", 7777},
			{"192.168.1.12", 7777},
		},
	}
}

func (m *Manager) AllocateServer() GameServer {

	server := m.servers[m.index]

	m.index++

	if m.index >= len(m.servers) {
		m.index = 0
	}

	return server
}

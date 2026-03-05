package party

import (
	"sync"

	"github.com/google/uuid"
)

type Party struct {
	ID      string     `json:"id"`
	Leader  string     `json:"leader"`
	Members []string   `json:"members"`
	Mutex   sync.Mutex `json:"-"`
}

func NewParty(leader string) *Party {

	id := uuid.New().String()

	return &Party{
		ID:      id,
		Leader:  leader,
		Members: []string{leader},
	}
}

func (p *Party) AddMember(player string) bool {

	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	for _, m := range p.Members {
		if m == player {
			return false
		}
	}

	p.Members = append(p.Members, player)
	return true
}

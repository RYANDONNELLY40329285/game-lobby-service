package matchmaking

import (
	"log"
	"sync"

	"github.com/ryandonnelly/game-lobby-service/internal/match"
	"github.com/ryandonnelly/game-lobby-service/internal/party"
)

type Queue struct {
	Parties []*party.Party
	Mutex   sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		Parties: []*party.Party{},
	}
}

func (q *Queue) JoinSolo(player string) *match.Match {

	p := &party.Party{
		ID:      player,
		Leader:  player,
		Members: []string{player},
	}

	return q.JoinParty(p)
}

func (q *Queue) JoinParty(p *party.Party) *match.Match {

	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	log.Println("Party joined matchmaking:", p.ID)

	q.Parties = append(q.Parties, p)

	const matchSize = 4

	players := []string{}
	usedParties := map[int]bool{}

	for i, party := range q.Parties {

		if len(players)+len(party.Members) > matchSize {
			continue
		}

		players = append(players, party.Members...)
		usedParties[i] = true

		if len(players) == matchSize {
			break
		}
	}

	if len(players) < matchSize {
		return nil
	}

	// create match using your match system
	m := match.NewMatch(players)

	// rebuild queue without used parties
	newQueue := []*party.Party{}

	for i, party := range q.Parties {
		if !usedParties[i] {
			newQueue = append(newQueue, party)
		}
	}

	q.Parties = newQueue

	return m
}

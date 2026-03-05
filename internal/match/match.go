package match

import "github.com/google/uuid"

type Match struct {
	ID     string   `json:"match_id"`
	TeamA  []string `json:"team_a"`
	TeamB  []string `json:"team_b"`
	Map    string   `json:"map"`
	Server string   `json:"server"`
	Status string   `json:"status"`
}

func NewMatch(players []string) *Match {

	half := len(players) / 2

	return &Match{
		ID:     uuid.New().String(),
		TeamA:  players[:half],
		TeamB:  players[half:],
		Map:    NextMap(),
		Server: AllocateServer(),
		Status: "pregame",
	}
}

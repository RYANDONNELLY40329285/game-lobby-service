package match

import "github.com/google/uuid"

type Match struct {
	ID       string   `json:"match_id"`
	TeamA    []string `json:"team_a"`
	TeamB    []string `json:"team_b"`
	Map      string   `json:"map"`
	ServerIP string   `json:"server_ip"`
	Port     int      `json:"port"`
	Status   string   `json:"status"`
}

func NewMatch(players []string, serverIP string, port int) *Match {

	half := len(players) / 2

	return &Match{
		ID:       uuid.New().String(),
		TeamA:    players[:half],
		TeamB:    players[half:],
		Map:      NextMap(),
		ServerIP: serverIP,
		Port:     port,
		Status:   "pregame",
	}
}

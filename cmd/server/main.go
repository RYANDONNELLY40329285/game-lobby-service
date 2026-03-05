package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ryandonnelly/game-lobby-service/internal/lobby"
	"github.com/ryandonnelly/game-lobby-service/internal/matchmaking"
	"github.com/ryandonnelly/game-lobby-service/internal/party"
	ws "github.com/ryandonnelly/game-lobby-service/internal/websocket"
)

var manager = lobby.NewManager()
var queue = matchmaking.NewQueue(manager)
var partyManager = party.NewManager()
var hub = ws.NewHub()

func createLobby(w http.ResponseWriter, r *http.Request) {

	host := r.URL.Query().Get("player")

	l := manager.CreateLobby(host)

	json.NewEncoder(w).Encode(l)
}

func joinLobby(w http.ResponseWriter, r *http.Request) {

	lobbyID := r.URL.Query().Get("lobby")
	player := r.URL.Query().Get("player")

	l, exists := manager.GetLobby(lobbyID)

	if !exists {
		http.Error(w, "Lobby not found", http.StatusNotFound)
		return
	}

	success := l.AddPlayer(player)

	if !success {
		http.Error(w, "Lobby full or player already joined", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(l)
}

func soloQueue(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Query().Get("player")

	l := queue.JoinSolo(player)

	if l == nil {
		w.Write([]byte("Searching for match..."))
		return
	}

	json.NewEncoder(w).Encode(l)
}
func createParty(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Query().Get("player")

	p := partyManager.CreateParty(player)

	json.NewEncoder(w).Encode(p)
}

func searchMatch(w http.ResponseWriter, r *http.Request) {

	partyID := r.URL.Query().Get("party")

	p, exists := partyManager.GetParty(partyID)

	if !exists {
		http.Error(w, "Party not found", http.StatusNotFound)
		return
	}

	l := queue.JoinParty(p)

	if l == nil {
		w.Write([]byte("Searching for match..."))
		return
	}

	json.NewEncoder(w).Encode(l)
}

func joinParty(w http.ResponseWriter, r *http.Request) {

	partyID := r.URL.Query().Get("party")
	player := r.URL.Query().Get("player")

	p, exists := partyManager.GetParty(partyID)

	if !exists {
		http.Error(w, "Party not found", http.StatusNotFound)
		return
	}

	success := p.AddMember(player)

	if !success {
		http.Error(w, "Player already in party", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/lobbies/create", createLobby)
	mux.HandleFunc("/lobbies/join", joinLobby)

	mux.HandleFunc("/matchmaking/solo", soloQueue)
	mux.HandleFunc("/matchmaking/search", searchMatch)

	mux.HandleFunc("/party/create", createParty)
	mux.HandleFunc("/party/join", joinParty)

	// WebSocket endpoint
	mux.HandleFunc("/ws", hub.HandleConnection)

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", mux)
}

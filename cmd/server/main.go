package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ryandonnelly/game-lobby-service/internal/gameserver"
	"github.com/ryandonnelly/game-lobby-service/internal/matchmaking"
	"github.com/ryandonnelly/game-lobby-service/internal/party"
	ws "github.com/ryandonnelly/game-lobby-service/internal/websocket"
)

var serverManager = gameserver.NewManager()
var queue = matchmaking.NewQueue(serverManager)
var partyManager = party.NewManager()
var hub = ws.NewHub()

func soloQueue(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Query().Get("player")

	match := queue.JoinSolo(player)

	if match == nil {
		w.Write([]byte("Searching for match..."))
		return
	}

	json.NewEncoder(w).Encode(match)
}

func createParty(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Query().Get("player")

	p := partyManager.CreateParty(player)

	json.NewEncoder(w).Encode(p)
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

func searchMatch(w http.ResponseWriter, r *http.Request) {

	partyID := r.URL.Query().Get("party")

	p, exists := partyManager.GetParty(partyID)

	if !exists {
		http.Error(w, "Party not found", http.StatusNotFound)
		return
	}

	match := queue.JoinParty(p)

	if match == nil {
		w.Write([]byte("Searching for match..."))
		return
	}

	json.NewEncoder(w).Encode(match)
}

func main() {

	mux := http.NewServeMux()

	// matchmaking
	mux.HandleFunc("/matchmaking/solo", soloQueue)
	mux.HandleFunc("/matchmaking/search", searchMatch)

	// party
	mux.HandleFunc("/party/create", createParty)
	mux.HandleFunc("/party/join", joinParty)

	// websocket
	mux.HandleFunc("/ws", hub.HandleConnection)

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", mux)
}

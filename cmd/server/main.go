package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ryandonnelly/game-lobby-service/internal/lobby"
)

var manager = lobby.NewManager()

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

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/lobbies/create", createLobby)
	mux.HandleFunc("/lobbies/join", joinLobby)

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", mux)
}

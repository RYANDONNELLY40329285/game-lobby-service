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

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/lobbies/create", createLobby)

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", mux)
}

package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn   *websocket.Conn
	Player string
	Lobby  string
}

type Hub struct {
	Lobbies map[string]map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		Lobbies: make(map[string]map[*Client]bool),
	}
}

func (h *Hub) BroadcastToLobby(lobby string, message []byte) {

	for client := range h.Lobbies[lobby] {

		err := client.Conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			log.Println(err)
			client.Conn.Close()
			delete(h.Lobbies[lobby], client)
		}
	}
}

func (h *Hub) HandleConnection(w http.ResponseWriter, r *http.Request) {

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	player := r.URL.Query().Get("player")
	lobby := r.URL.Query().Get("lobby")

	client := &Client{
		Conn:   conn,
		Player: player,
		Lobby:  lobby,
	}

	// create lobby room if it doesn't exist
	if h.Lobbies[lobby] == nil {
		h.Lobbies[lobby] = make(map[*Client]bool)
	}

	h.Lobbies[lobby][client] = true

	log.Println(player, "connected to lobby", lobby)

	for {

		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Println(player, "disconnected")

			delete(h.Lobbies[lobby], client)
			conn.Close()
			break
		}

		log.Println("Message from", player, ":", string(msg))

		h.BroadcastToLobby(lobby, msg)
	}
}

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
	Clients map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		Clients: make(map[*Client]bool),
	}
}

func (h *Hub) Broadcast(message []byte) {

	for client := range h.Clients {

		err := client.Conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			log.Println(err)
			client.Conn.Close()
			delete(h.Clients, client)
		}
	}
}

func (h *Hub) HandleConnection(w http.ResponseWriter, r *http.Request) {

	conn, err := Upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		Conn: conn,
	}

	h.Clients[client] = true

	log.Println("Client connected")

	for {

		_, msg, err := conn.ReadMessage()

		if err != nil {
			delete(h.Clients, client)
			conn.Close()
			break
		}

		log.Println("Received:", string(msg))

		h.Broadcast(msg)
	}
}

# game-lobby-service
Multiplayer Game Lobby & Matchmaking Service

A backend service for multiplayer game matchmaking built in Go, supporting party systems, matchmaking queues, game server allocation, and WebSocket-based real-time communication.

This project simulates the backend architecture used by modern multiplayer games such as Call of Duty, Valorant, Apex Legends, and Fortnite.

Features
Party System

Players can create and manage parties before matchmaking.

Create Party
Join Party
Party Leader controls matchmaking

Example party:

Party
------
Ryan (Leader)
John
Alice
Bob
Solo & Party Matchmaking

Players can queue:

Solo
Duo
Party

The matchmaking queue groups players into matches while ensuring party integrity (party members stay together).

Example queue:

Queue
------
Party (3)
Solo
Party (2)
Solo

Match result:

Team A
Ryan
John
Alice

Team B
Mike
Sarah
Tom
Match Creation

When enough players join the queue, the system generates a match object containing:

match_id
teams
map
server_ip
port
status

Example response:

{
 "match_id": "b41c1c3c-34b2-4f6c-a98d-3b3a49c9e2f3",
 "team_a": ["ryan","john"],
 "team_b": ["alice","bob"],
 "map": "Rust",
 "server_ip": "192.168.1.10",
 "port": 7777,
 "status": "pregame"
}
Map Rotation

Matches automatically rotate maps from a predefined pool.

Example rotation:

Nuketown
Rust
Shipment
Terminal
Game Server Allocation

The system assigns matches to a simulated game server pool.

Example servers:

192.168.1.10:7777
192.168.1.11:7777
192.168.1.12:7777

This mimics real multiplayer infrastructure where matchmaking assigns players to running game servers.

WebSocket Lobby Communication

Players connect to a WebSocket server for real-time lobby updates.

Example:

ws://localhost:8080/ws

This allows:

Lobby chat
Player join notifications
Game start events
Architecture
Client
  │
API
  │
Party Service
  │
Matchmaking Queue
  │
Match Service
  │
Game Server Manager
  │
Game Server

Components:

internal/
 ├ party
 ├ matchmaking
 ├ match
 ├ gameserver
 └ websocket
API Endpoints
Solo Matchmaking
GET /matchmaking/solo?player=ryan
Create Party
GET /party/create?player=ryan
Join Party
GET /party/join?party=PARTY_ID&player=john
Party Matchmaking
GET /matchmaking/search?party=PARTY_ID
WebSocket Lobby
ws://localhost:8080/ws
Running the Server

Start the backend service:

go run ./cmd/server

Server runs on:

http://localhost:8080
Testing Matchmaking

Example solo queue test:

curl "http://localhost:8080/matchmaking/solo?player=ryan"
curl "http://localhost:8080/matchmaking/solo?player=john"
curl "http://localhost:8080/matchmaking/solo?player=alice"
curl "http://localhost:8080/matchmaking/solo?player=bob"

Once enough players join, a match will be created automatically.

Simulation Script

A Python script is included to simulate multiple players joining matchmaking simultaneously.

Example simulation:

2 simultaneous lobbies
party matchmaking
solo matchmaking
server allocation

Run:

python simulate_lobbies.py
Technologies Used
Go
WebSockets
HTTP APIs
Concurrent Matchmaking Queues
Python (load simulation)
package socket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	// Clients registered in socket
	Clients map[uuid.UUID]*websocket.Conn
	// Broadcast saved in socket
	Broadcast chan Message
)

// ConfigSocket configure chat socket
func ConfigSocket() {
	Clients = make(map[uuid.UUID]*websocket.Conn)
	Broadcast = make(chan Message)
}

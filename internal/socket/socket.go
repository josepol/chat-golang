package socket

import (
	"github.com/gorilla/websocket"
)

var (
	// Clients registered in socket
	Clients map[*websocket.Conn]bool
	// Broadcast saved in socket
	Broadcast chan Message
)

// ConfigSocket configure chat socket
func ConfigSocket() {
	Clients = make(map[*websocket.Conn]bool)
	Broadcast = make(chan Message)
}

package chat

import (
	"api/internal/socket"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func socketRoute(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting chat socket...")

	socket.ConfigSocket()

	go handleMessages()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// defer conn.Close()

	socket.Clients[conn] = true

	for {
		var message socket.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("Socket error %v", err)
			delete(socket.Clients, conn)
		}
		socket.Broadcast <- message
	}

}

func handleMessages() {
	for {
		msg := <-socket.Broadcast
		for client := range socket.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Socket error %v", err)
				client.Close()
				delete(socket.Clients, client)
			}
		}
	}
}

package chat

import (
	"api/internal/socket"
	"encoding/json"
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

	/*if r.Header.Get("Origin") != "http://"+r.Host {
		log.Print(r.Header.Get("Origin"))
		log.Print(r.Host)
		http.Error(w, "Origin not allowed", 403)
		return
	}*/

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	socket.Clients[conn] = true

	for {
		var message socket.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("Socket error %v", err)
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			// delete(socket.Clients, conn)
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

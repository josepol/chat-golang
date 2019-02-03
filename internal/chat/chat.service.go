package chat

import (
	"api/internal/socket"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func socketRoute(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting chat socket...")

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

	var id uuid.UUID
	id, _ = uuid.NewUUID()

	socket.Clients[id] = conn
	log.Print(socket.Clients)

	defer conn.Close()

	for {
		var message socket.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("ReadJSON socket error %v", err)
			delete(socket.Clients, id)
		}
		socket.Broadcast <- message
	}

}

func handleMessages() {
	for {
		msg := <-socket.Broadcast
		log.Print(len(socket.Clients))
		for key, clientConn := range socket.Clients {
			err := clientConn.WriteJSON(msg)
			if err != nil {
				log.Printf("WriteJSON socket error %v", err)
				clientConn.Close()
				delete(socket.Clients, key)
			}
		}
	}
}

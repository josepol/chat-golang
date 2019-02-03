package main

import (
	"api/internal/route"
	"api/internal/socket"
)

func main() {
	socket.ConfigSocket()
	route.ConfigAPIRouting()
}

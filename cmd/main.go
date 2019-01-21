package main

import (
	"api/internal/database"
	"api/internal/route"
)

func main() {
	database.ConfigDatabase()
	route.Config()
}

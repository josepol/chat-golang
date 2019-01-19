package main

import (
	route "api/internal"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	setupMySQL()
	route.Config()
}

func setupMySQL() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	/*uuid, _ := uuid.NewRandom()

	insert, err := db.Query("INSERT INTO auth VALUES (?, 'asdasdasd', '123123')", uuid)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()*/
}

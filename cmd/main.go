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

	insert, err := db.Query("INSERT INTO auth VALUES (1, 'josepol', '123123')")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

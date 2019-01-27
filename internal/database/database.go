package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DB mysql connection
var (
	DB              *sql.DB
	DuplicatedError uint16 = 1062
)

// ConfigDatabase Connect to MySQL Database
func OpenDBConnection() {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err.Error())
	}
}

// CloseDatabase Close database connection
func CloseDBConnection() {
	defer DB.Close()
}

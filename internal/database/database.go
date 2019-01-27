package database

import (
	"database/sql"
)

// DB mysql connection
// DuplicatedError SQL duplicated row
var (
	DB              *sql.DB
	DuplicatedError uint16 = 1062
)

// OpenDBConnection Connect to MySQL Database
func OpenDBConnection() {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err.Error())
	}
}

// CloseDBConnection Close database connection
func CloseDBConnection() {
	defer DB.Close()
}

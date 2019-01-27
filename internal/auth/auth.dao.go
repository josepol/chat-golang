package auth

import (
	"api/internal/database"
	"database/sql"
)

func init() {}

func loginDao(auth Auth) (*sql.Rows, error) {
	query, err := database.DB.Query("SELECT id, username, password FROM auth WHERE username = ? AND password = ?", auth.Username, auth.Password)
	return query, err
}

func registerDao(auth Auth) (*sql.Rows, error) {
	insert, err := database.DB.Query("INSERT INTO auth (id, username, password) VALUES (?, ?, ?)", auth.ID, auth.Username, auth.Password)
	return insert, err
}

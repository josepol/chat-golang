package auth

import (
	"api/internal/database"
	"database/sql"
)

func loginDao(auth Auth) (*sql.Rows, error) {
	query, err := database.DB.Query("SELECT id, username, password FROM auth WHERE username = ?", auth.Username)
	return query, err
}

func registerDao(auth Auth) (*sql.Rows, error) {
	insert, err := database.DB.Query("INSERT INTO auth (id, username, password) VALUES (?, ?, ?)", auth.ID, auth.Username, auth.Password)
	return insert, err
}

func getUsername(id string) (*sql.Rows, error) {
	query, err := database.DB.Query("SELECT username FROM auth WHERE id = ?", id)
	return query, err
}

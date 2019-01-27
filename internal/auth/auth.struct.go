package auth

import (
	"github.com/google/uuid"
)

// Auth struct
type Auth struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

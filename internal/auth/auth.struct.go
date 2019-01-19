package auth

// Auth struct
type Auth struct {
	ID       [16]byte `json:"id"`
	Username string   `json: "username"`
	Password string   `json: "password"`
}

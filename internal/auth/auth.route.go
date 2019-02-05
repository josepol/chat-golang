package auth

import (
	"github.com/gorilla/mux"
)

func init() {
}

// Config Auth routing
func Config(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/login", login).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/register", register).Methods("POST")
	router.HandleFunc("/auth/test", withMiddleware(test)).Methods("POST")
	return router
}

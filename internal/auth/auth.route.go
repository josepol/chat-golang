package auth

import (
	"github.com/gorilla/mux"
)

func init() {
}

// Config routing
func Config(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/login", login).Methods("POST")
	router.HandleFunc("/auth/register", register).Methods("POST")
	return router
}

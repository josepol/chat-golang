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
	router.HandleFunc("/auth/test", withMiddleware(test)).Methods("POST")
	return router
}

package chat

import (
	"github.com/gorilla/mux"
)

// Config Chat routing
func Config(router *mux.Router) *mux.Router {
	router.HandleFunc("/chat/socket", socketRoute)
	return router
}

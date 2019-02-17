package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	auth "api/internal/auth"
	"api/internal/chat"
)

// ConfigAPIRouting API routing
func ConfigAPIRouting() {
	router := mux.NewRouter()
	router = auth.Config(router)
	router = chat.Config(router)
	router.HandleFunc("/auth/test/{msg}", testConnection).Methods("GET")

	fs := http.FileServer(http.Dir("../../public"))
	http.Handle("/", fs)

	handler := cors.Default().Handler(router)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})

	// Insert the middleware
	handler = c.Handler(handler)

	log.Fatal(http.ListenAndServe(":3001", handler))
}

func testConnection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

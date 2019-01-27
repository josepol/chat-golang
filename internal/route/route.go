package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	auth "api/internal/auth"
)

// Config API routing
func Config() {
	router := mux.NewRouter()
	router = auth.Config(router)
	router.HandleFunc("/auth/test/{msg}", testConnection).Methods("GET")

	headersOk, originOk, methodOk := headers()
	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(originOk, headersOk, methodOk)(router)))
}

func headers() (handlers.CORSOption, handlers.CORSOption, handlers.CORSOption) {
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	return headersOk, originsOk, methodsOk
}

func testConnection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

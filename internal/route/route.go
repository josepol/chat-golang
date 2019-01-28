package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

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

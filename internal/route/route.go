package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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

	// headersOk, originOk, methodOk := headers()
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3001" /*handlers.CORS(originOk, headersOk, methodOk)*/, handler))
}

func headers() (handlers.CORSOption, handlers.CORSOption, handlers.CORSOption) {
	headersOk := handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return headersOk, originsOk, methodsOk
}

func testConnection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

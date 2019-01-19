package auth

import (
	"encoding/json"
	"net/http"
)

func init() {}

func login(w http.ResponseWriter, r *http.Request) {
	loginDao()
	json.NewEncoder(w).Encode("Auth login!")
}

func register(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	message := vars.Get("msg")

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

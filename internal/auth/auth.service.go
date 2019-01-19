package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {}

func login(w http.ResponseWriter, r *http.Request) {
	loginDao()
	json.NewEncoder(w).Encode("Auth login!")
}

func register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var auth Auth
	err := decoder.Decode(&auth)
	if err != nil {
		panic(err)
	}

	fmt.Print(auth)

	json.NewEncoder(w).Encode(map[string]string{"username": auth.Username})
}

package auth

import (
	"api/internal/database"
	routestruct "api/internal/route/struct"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func init() {}

func login(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Auth login!")
}

func register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var auth Auth
	err := decoder.Decode(&auth)

	if err != nil {
		panic(err)
	}

	auth.ID, err = uuid.NewUUID()

	if err != nil {
		panic(err.Error())
	}

	dbStatus, err := registerDao(auth)

	if err != nil {
		if err.(*mysql.MySQLError).Number == database.DuplicatedError {
			var errorMessage = routestruct.ErrorMessage{Status: "ERR", Message: "DUPLICATED"}
			json.NewEncoder(w).Encode(errorMessage)
		} else {
			json.NewEncoder(w).Encode(err.Error())
		}
		return
	}

	fmt.Print(dbStatus)

	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})

}

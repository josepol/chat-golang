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

var auth Auth
var statusResponse routestruct.StatusResponse

func login(w http.ResponseWriter, r *http.Request) {

	authFound := Auth{}

	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&auth)

	if decodeErr != nil {
		responseJSON(w, getGenericError())
	}

	rows, loginErr := loginDao(auth)

	if loginErr != nil {
		responseJSON(w, getGenericError())
	}

	for rows.Next() {
		var id uuid.UUID
		var username, password string
		selectedUserErr := rows.Scan(&id, &username, &password)
		if selectedUserErr != nil {
			responseJSON(w, getGenericError())
		}
		authFound.ID = id
		authFound.Username = username
		authFound.Password = password
	}

	statusResponse = routestruct.StatusResponse{Status: "02", Message: fmt.Sprintf("%b", authFound.ID)}
	responseJSON(w, statusResponse)
}

func register(w http.ResponseWriter, r *http.Request) {

	var uuidErr error
	var registerErr error

	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&auth)

	auth.ID, uuidErr = uuid.NewUUID()

	if uuidErr != nil || decodeErr != nil {
		statusResponse = routestruct.StatusResponse{Status: "02", Message: "There has been an error"}
		responseJSON(w, statusResponse)
	}

	_, registerErr = registerDao(auth)

	if registerErr != nil {
		if registerErr.(*mysql.MySQLError).Number == database.DuplicatedError {
			statusResponse = routestruct.StatusResponse{Status: "01", Message: "User already registered"}
		} else {
			statusResponse = routestruct.StatusResponse{Status: "02", Message: "There has been an error"}
		}
	} else {
		statusResponse = routestruct.StatusResponse{Status: "00", Message: "The registering operation has been succeded"}
	}

	responseJSON(w, statusResponse)

}

func getGenericError() routestruct.StatusResponse {
	return routestruct.StatusResponse{Status: "02", Message: "There has been an error"}
}

func responseJSON(w http.ResponseWriter, statusResponse routestruct.StatusResponse) {
	json.NewEncoder(w).Encode(statusResponse)
}

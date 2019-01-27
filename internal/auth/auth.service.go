package auth

import (
	"api/internal/database"
	routestruct "api/internal/route/struct"
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var auth Auth
var statusResponse routestruct.StatusResponse

func login(w http.ResponseWriter, r *http.Request) {

	database.OpenDBConnection()

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

	database.CloseDBConnection()
	json.NewEncoder(w).Encode(authFound)
}

func register(w http.ResponseWriter, r *http.Request) {

	database.OpenDBConnection()

	var uuidErr error
	var registerErr error

	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&auth)
	auth.ID, uuidErr = uuid.NewUUID()

	if uuidErr != nil || decodeErr != nil {
		responseJSON(w, getGenericError())
	}

	_, registerErr = registerDao(auth)

	if registerErr != nil {
		if registerErr.(*mysql.MySQLError).Number == database.DuplicatedError {
			statusResponse = routestruct.StatusResponse{Status: "01", Message: "User already registered"}
		} else {
			statusResponse = getGenericError()
		}
	} else {
		statusResponse = getGenericSuccess()
	}

	database.CloseDBConnection()
	responseJSON(w, statusResponse)

}

func getGenericError() routestruct.StatusResponse {
	return routestruct.StatusResponse{Status: "02", Message: "There has been an error"}
}

func getGenericSuccess() routestruct.StatusResponse {
	return routestruct.StatusResponse{Status: "02", Message: "Operation has been succeded"}
}

func responseJSON(w http.ResponseWriter, statusResponse routestruct.StatusResponse) {
	json.NewEncoder(w).Encode(statusResponse)
}

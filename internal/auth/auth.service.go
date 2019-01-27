package auth

import (
	"api/internal/database"
	model "api/internal/model"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var auth Auth
var statusResponse model.StatusResponse

func login(w http.ResponseWriter, r *http.Request) {

	database.OpenDBConnection()

	authFound := Auth{}
	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&auth)

	if decodeErr != nil {
		responseJSON(w, getGenericError())
		return
	}

	rows, loginErr := loginDao(auth)

	if loginErr != nil {
		responseJSON(w, getGenericError())
		return
	}

	for rows.Next() {
		var id uuid.UUID
		var username, password string
		selectedUserErr := rows.Scan(&id, &username, &password)
		if selectedUserErr != nil {
			responseJSON(w, getGenericError())
			return
		}
		authFound.ID = id
		authFound.Username = username
		authFound.Password = password
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(authFound.Password), []byte(auth.Password))

	if passwordErr != nil {
		statusResponse = model.StatusResponse{Status: "03", Message: "Login failed"}
		responseJSON(w, statusResponse)
		return
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
		return
	}

	hashedPassword, passwordCryptErr := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)

	auth.Password = string(hashedPassword)

	if passwordCryptErr != nil {
		responseJSON(w, getGenericError())
		return
	}

	_, registerErr = registerDao(auth)

	if registerErr != nil {
		if registerErr.(*mysql.MySQLError).Number == database.DuplicatedError {
			statusResponse = model.StatusResponse{Status: "01", Message: "User already registered"}
		} else {
			statusResponse = getGenericError()
		}
	} else {
		statusResponse = getGenericSuccess()
	}

	database.CloseDBConnection()
	responseJSON(w, statusResponse)

}

func getGenericError() model.StatusResponse {
	return model.StatusResponse{Status: "02", Message: "There has been an error"}
}

func getGenericSuccess() model.StatusResponse {
	return model.StatusResponse{Status: "02", Message: "Operation has been succeded"}
}

func responseJSON(w http.ResponseWriter, statusResponse model.StatusResponse) {
	json.NewEncoder(w).Encode(statusResponse)
}

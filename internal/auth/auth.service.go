package auth

import (
	"api/internal/database"
	model "api/internal/model"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
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
		statusResponse = model.StatusResponse{Status: "3", Message: "Login failed"}
		responseJSON(w, statusResponse)
		return
	}

	jwt := generateJWT()

	database.CloseDBConnection()
	statusResponse = model.StatusResponse{Status: "0", Message: jwt}
	responseJSON(w, statusResponse)
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
			statusResponse = model.StatusResponse{Status: "1", Message: "User already registered"}
		} else {
			statusResponse = getGenericError()
		}
		responseJSON(w, statusResponse)
		return
	}

	jwt := generateJWT()

	database.CloseDBConnection()
	responseJSON(w, model.StatusResponse{Status: "0", Message: jwt})

}

func test(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.StatusResponse{Status: "00", Message: "OK"})
}

func getGenericError() model.StatusResponse {
	return model.StatusResponse{Status: "2", Message: "There has been an error"}
}

func getGenericSuccess() model.StatusResponse {
	return model.StatusResponse{Status: "2", Message: "Operation has been succeded"}
}

func responseJSON(w http.ResponseWriter, statusResponse model.StatusResponse) {
	json.NewEncoder(w).Encode(statusResponse)
}

func generateJWT() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = auth.ID
	token.Claims = claims
	tokenString, jwtSignedErr := token.SignedString([]byte("privateKey123"))

	if jwtSignedErr != nil {
		panic(jwtSignedErr)
	}

	return tokenString
}

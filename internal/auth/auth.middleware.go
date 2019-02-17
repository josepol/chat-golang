package auth

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

func withMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationToken := r.Header.Get("Authorization")
		log.Printf("Logged connection token %s", authorizationToken)

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(authorizationToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("privateKey123"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Print(claims["id"])
			context.Set(r, "userID", claims["id"])
		} else {
			fmt.Println(err)
		}

		next.ServeHTTP(w, r)
	}
}

package auth

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
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
			fmt.Println(claims["id"])
			log.Printf("User id is %s", claims["id"])
		} else {
			fmt.Println(err)
		}

		next.ServeHTTP(w, r)
	}
}

package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("asdasfas"), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func runJwtAuthentication() {
	// Check https://hackernoon.com/creating-a-middleware-in-golang-for-jwt-based-authentication-cx3f32z8 for full story
	/*
		While most forms of token authentication requires a database read to verify the token belongs to an
		active authenticated user, when using JWTs, if the JWT can be decoded successfully, that itself
		guarantees it is a valid token since it has a signature field that will become invalid if any data
		in the token is corrupted or manipulated. You can also decide what data to encode in the JWT body,
		which means on successfully decoding a JWT you can also get useful data, such as a user's username
		or email.

		The scope of this article is limited to creating a middleware in Golang to check the validity of a
		JWT in an incoming request.
	*/
	http.Handle("/ping", middleware(http.HandlerFunc(pong)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
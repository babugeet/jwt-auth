package utils

import (
	"context"
	"fmt"
	"jwt-auth/mocks"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var UserMap map[string]string

func Struct2Map() map[string]string {
	mockss := mocks.Users
	// fmt.Println("this is mocks")
	// fmt.Println(mockss)
	// fmt.Println("caled struct2Map")
	m := make(map[string]string)
	for _, j := range mockss {
		fmt.Println(j.Username)
		fmt.Println(j.Password)
		m[j.Username] = j.Password
	}
	// fmt.Println("This is usermap")
	// fmt.Println(m)
	return m
}

var jwtKey = []byte("secret_key")

// AuthMiddleware is a middleware that checks for a valid JWT token.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println(tokenString)
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		fmt.Println(err)
		if err != nil || !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Store the claims in the context (optional, if you want to access it later)
		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", claims.Username)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}

// Expect the token from Authorization header
// authHeader := r.Header.Get("Authorization")
// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 	w.WriteHeader(http.StatusUnauthorized)
// 	return
// }
// tokenString := strings.TrimPrefix(authHeader, "Bearer ")

// claims := &utils.Claims{}
// tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
// 	return jwtKey, nil
// })
// if err != nil || !tkn.Valid {
// 	w.WriteHeader(http.StatusUnauthorized)
// 	return
// }

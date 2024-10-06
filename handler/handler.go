package handler

import (
	"encoding/json"
	"fmt"
	"jwt-auth/mocks"
	"jwt-auth/models"
	"jwt-auth/utils"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reached login")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pass_val, ok := utils.Struct2Map()[user.Username]
	if !ok || pass_val != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime4Token := time.Now().Add(time.Minute * 5)
	claims := utils.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime4Token.Unix(),
		},
	}
	newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newtoken.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the token as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

// Refresh handler
func Refresh(w http.ResponseWriter, r *http.Request) {
	// Expect the token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &utils.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a new token with a refreshed expiration time
	expirationTime4Token := time.Now().Add(time.Minute * 5)
	claims.ExpiresAt = expirationTime4Token.Unix()
	newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = newtoken.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the refreshed token in response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

// Signup handler
func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, ok := utils.Struct2Map()[user.Username]
	if ok {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("User already exists, please login"))
		return
	}
	mocks.Users = append(mocks.Users, user)
	w.Write([]byte("User profile created successfully"))
}

// CheckAuth handler
func CheckAuth(w http.ResponseWriter, r *http.Request) {
	// Expect the token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &utils.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
}

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	// Expect the token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &utils.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
}

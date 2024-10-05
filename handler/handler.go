package handler

import (
	"encoding/json"
	"fmt"
	"jwt-auth/models"
	"jwt-auth/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my-secret")

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pass_val, ok := utils.Struct2Map()[user.UserName]
	if !ok || pass_val != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		expirationTime4Token := time.Now().Add(time.Minute * 5)
		claimss := utils.Claims{
			Username: user.UserName,
			StandardClaims: jwt.StandardClaims{
				Audience:  "",
				ExpiresAt: expirationTime4Token.Unix(),
			},
		}
		newtoken := jwt.NewWithClaims(jwt.SigningMethodES256, claimss)
		tokenString, err := newtoken.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: tokenString,

			Expires: expirationTime4Token,
		})
	}

}

func Refresh(w http.ResponseWriter, r *http.Request) {

}

func Home(w http.ResponseWriter, r *http.Request) {

	cokkie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenString := cokkie.Value
	claims := &utils.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
}

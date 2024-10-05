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

var jwtKey = []byte("secret_key")

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reached login")
	var user models.User
	// fmt.Println(r.Body.Read()
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user.Password, user.Username)
	fmt.Println(utils.Struct2Map())
	pass_val, ok := utils.Struct2Map()[user.Username]

	fmt.Println(pass_val)
	if !ok || pass_val != user.Password {
		// fmt.Println("Reached here")
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		// fmt.Println("Reached here")
		expirationTime4Token := time.Now().Add(time.Minute * 5)
		claimss := utils.Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				Audience:  "",
				ExpiresAt: expirationTime4Token.Unix(),
			},
		}
		newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimss)
		tokenString, err := newtoken.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		fmt.Println(tokenString)
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

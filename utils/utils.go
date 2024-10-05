package utils

import (
	"jwt-auth/mocks"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var UserMap map[string]string

func Struct2Map() map[string]string {
	mockss := mocks.Users
	for _, j := range mockss {
		UserMap[j.UserName] = j.Password
	}
	return UserMap
}

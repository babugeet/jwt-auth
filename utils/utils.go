package utils

import (
	"fmt"
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

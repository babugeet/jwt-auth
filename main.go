package main

import (
	"jwt-auth/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/home", handler.Home)
	http.HandleFunc("/refresh", handler.Refresh)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
